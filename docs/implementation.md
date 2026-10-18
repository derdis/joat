# OAuth dashboard implementation plan

## Goal and order

Build a working, observable OAuth-style demo before building the designed dashboard.

The first usable interface is deliberately plain: two continuously updating log columns, **server traffic** and **your activity**, plus Login, Make protected request, and Logout buttons. A visitor can open it anonymously, identify their own requests, click Login without entering credentials, and return to their completed login trace.

Implementation order:

1. Replace joat's Gin scaffold with standard-library HTTP handlers.
2. Make notauth reliable and configurable.
3. Add the credential-free demo login.
4. Integrate notauth's two services into joat's request path.
5. Establish anonymous visitor and login-flow correlation.
6. Collect structured auth events in the observation backend.
7. Ingest Nginx access logs.
8. Deliver snapshots and continuous SSE streams.
9. Replace the Next.js scaffold with React + Vite and add the temporary two-column interface.
10. Validate the end-to-end milestone, then implement JWT/JWKS validation and its observable stages.
11. Build the final dashboard and visual design.

This document is a plan. Creating it does not authorize implementation, deployment, or publication.

## Agreed technology decisions

- **Backend: keep the Go service, replace Gin with `net/http`.** Use `http.ServeMux`, ordinary `http.Handler`/`http.HandlerFunc` handlers, `encoding/json`, request contexts, and explicit response flushing. Keep event storage/subscriptions independent of HTTP. Add request logging and panic recovery deliberately when removing `gin.Default()`.
- **Frontend: React + Vite + TypeScript + Tailwind CSS.** Replace Next.js rather than keeping its App Router, server components, or server runtime. React handles the eventual interactive timeline and inspectors; Vite provides development and production builds.
- **Production frontend: static build artifacts served by Nginx.** Node/Bun can be used during the build, but no Next.js, Vite development/preview, or Node/Bun frontend server runs in production.
- **Auth, sessions, APIs, and SSE remain in Go.** The browser fetches a snapshot and subscribes to the Go event stream on the same public origin. Framework migration does not move token exchange or session ownership into JavaScript.
- **Design is independent of the framework.** Retain the option of Vercel-like typography and components within the existing design direction. Final styling and any component library choices come after the working two-column interface.

## Existing scaffold to migrate

These are the pre-migration files and limitations, not the chosen stack. The target is Nginx + Go `net/http` services + a static React/Vite frontend.

- Replace the existing Gin backend and Next.js frontend scaffolds configured in Compose with the target stack. Retain Nginx as the public proxy and static file server.
- Extend `nginx/nginx.conf`'s existing JSON access logs with request IDs and collection. Replace its current frontend-server proxy route with the static-serving route described in step 4; retain the text response at `/` until the final landing-page decision.
- Replace `backend/main.go`'s placeholder `/api` response and single-event `/events` handler with the planned standard-library handlers and continuous stream.
- Replace the template at `frontend/app/page.tsx` with the Vite entry point and React components in step 9. The old path is migration input, not the future frontend entry point.
- `/home/derdis/dev/notauth` contains independently runnable authserver and middleware programs, shared types, and a Dockerfile capable of building either program.
- notauth already implements authorization-code issuance/redemption, provider sessions, app sessions, and access/refresh token handling. Its access tokens are opaque strings. `/jwks` is a stub; signed JWTs and OIDC behavior must not be claimed yet.
- The middleware currently serves a text response instead of forwarding a protected request to joat.
- Both repositories have existing local changes. Preserve them and review their diffs before implementing this plan.

## Target service responsibilities

| Service | Responsibility |
| --- | --- |
| Nginx | Public TLS termination, route selection, static frontend serving, trusted request IDs, JSON access logging |
| notauth authserver | Demo identity, provider session, code issuance/redemption, token issuance |
| notauth middleware | Anonymous visitor context, login transactions, callback, app sessions, refresh, public/protected proxying |
| joat backend (`net/http`) | Protected demo handler, event ingestion, bounded history, snapshots, SSE |
| Access-log collector | Read Nginx JSON lines and submit sanitized traffic events |
| React frontend (Vite build) | Static browser application: temporary log interface first; designed dashboard last |

Keep notauth as the source of its auth code. Build its two binaries from a pinned sibling checkout during local development; make that source revision available explicitly to the VPS build/deployment workflow. Do not depend on an unexplained developer-machine path on the server or maintain duplicate copies of the Go source.

The middleware becomes a small application gateway. Public means no authenticated app session is required; it can still establish an anonymous visitor context. Observation endpoints must remain available before login.

The frontend row represents a build artifact, not a long-running application service. Preserve visitor creation on the first dashboard request by routing public dashboard requests through the gateway to a private static-only Nginx listener. That listener serves the Vite build directly and never routes back into the gateway. It is not published to the host, so the public route remains the visitor-aware entry point.

## 1. Replace the Gin scaffold

Primary files: `backend/main.go`, `backend/go.mod`, and `backend/go.sum`.

- [ ] Replace Gin routing and context handlers with an explicit `http.ServeMux` and standard `http.ResponseWriter`/`*http.Request` handlers.
- [ ] Replace `gin.H`/`c.JSON` with typed response values and `encoding/json`, setting content types and status codes explicitly.
- [ ] Keep only small health/demo placeholders until the observation endpoints are built. Remove the single-event SSE placeholder rather than treating it as a functioning stream.
- [ ] Add simple request logging and panic recovery around the handlers, preserving request IDs and redacting sensitive data.
- [ ] Configure an explicit `http.Server`, sensible request limits/timeouts, and graceful shutdown. Do not apply a short global write timeout that cuts off long-lived SSE responses.
- [ ] Remove Gin imports and run `go mod tidy` to remove dependencies no longer needed. Continue using established libraries for JWT/cryptographic functionality when those steps arrive.

Completion check: the backend builds and serves its health/demo endpoints without Gin. It uses the same handler primitives as notauth, and its HTTP shell does not dictate event-store design.

## 2. Stabilize and configure notauth

Primary files: `notauth/cmd/authserver/main.go`, `notauth/cmd/middleware/main.go`, and `notauth/internal/shared/helpers.go`.

- [ ] Replace hardcoded localhost URLs and client configuration with validated configuration: listen address, public auth URL, public callback URL, internal token URL, frontend/backend upstream URLs, client ID, and cookie settings.
- [ ] Separate public browser URLs from private service URLs. A container must call `authserver:9000`, not its own localhost; the browser must receive a public HTTPS URL.
- [ ] Fix the middleware's missing-cookie path: it must not dereference a missing cookie when another app session already exists.
- [ ] Protect mutable maps with synchronization; add expiration and cleanup for pending states, codes, sessions, and refresh tokens. Apply the same rules to visitor storage when it is introduced in step 5.
- [ ] Bind pending state to the initiating browser using a short-lived opaque transaction cookie and an expiry. Validate and consume it during callback handling; reject a callback initiated by a different browser. Link this transaction to the anonymous visitor context in step 5.
- [ ] Validate each authorization code against its registered client and redirect URI, and consume it atomically once. Derive the authenticated user from the provider session/code, not a browser-supplied username.
- [ ] Bind refresh-token redemption to the correct client and session. Rotate refresh tokens on redemption and invalidate the previous value.
- [ ] Check token-exchange status codes and JSON errors, close response bodies, apply timeouts, and reject incomplete token responses before creating an app session.
- [ ] Use consistent explicit cookie paths and lifetime rules. Public HTTPS cookies use Secure, HttpOnly, and an appropriate explicit SameSite setting; localhost development gets a separate configuration.
- [ ] Replace secret-bearing console output with structured, redacted diagnostics. Current output includes token values and callback URLs containing codes.
- [ ] Bring the authorization-code exchange toward the standard parameter shapes (`code`, `grant_type`, form-encoded token requests) and add PKCE S256 before public deployment. Update both programs together.

Completion check: two independent browser sessions can complete login and refresh without crashes or crossed identities. Invalid, expired, replayed, or incorrectly bound callbacks/codes/token requests fail explicitly and do not create sessions.

## 3. Add a credential-free demo login

- [ ] Add the middleware's `POST /auth/start` handler to create the browser-bound transaction and begin the authorization redirect. Public Nginx routing for it is added in step 4.
- [ ] Add an explicit demo mode to the authserver. Following the dashboard's login action, it creates a disposable pseudonymous identity without showing a credential form.
- [ ] Reuse a valid provider session where appropriate; create a new demo identity for a genuinely new provider session.
- [ ] Preserve the real sequence: authorize request, provider session, code issuance, callback, server-side code exchange, app session, redirect back.
- [ ] Remove the callback's dependency on a `username` query parameter. The redeemed code determines identity.
- [ ] Change the callback's final destination from the text `/app` page to `/dashboard/`.
- [ ] Allow only registered callbacks and a fixed/allowlisted post-login dashboard destination.
- [ ] Implement logout that expires the app cookie and revokes its session/refresh token. Provide an explicit full demo reset that also clears the provider session when a fresh identity is wanted.

Completion check: following the login-start endpoint in a fresh browser completes the redirect/code/session sequence without a credential form and redirects to `/dashboard/`. The temporary static destination is wired in step 4; the working log interface arrives in step 9. This establishes a demo identity, not verified human identity.

## 4. Wire Compose and the gateway routes

Primary joat files: `docker-compose.yaml`, `nginx/nginx.conf`, a new `nginx/Dockerfile` for the static production image, and `backend/main.go`. The Vite configuration/build is added in step 9; gateway changes remain in notauth.

- [ ] Add authserver and middleware services to joat's private Compose network. Both middleware and backend can listen on port 8080 inside separate containers; publish only Nginx's public ports.
- [ ] Configure the middleware to proxy public dashboard/observation requests and authenticated demo API requests to their respective upstreams.
- [ ] Prepare static frontend serving under `/dashboard/`, with `/dashboard` canonicalized to that path. Public dashboard requests go through the gateway, then to a private static-only Nginx listener; ensure this cannot create a proxy loop. Use a minimal static probe page until step 9 supplies the Vite build. Keep the existing root response until the final landing-page decision.
- [ ] Serve the built frontend's entry HTML for dashboard navigation routes, while missing asset paths return 404 rather than HTML. Keep visitor-specific responses and cookie-setting HTML uncached; immutable built assets may use long-lived caching.
- [ ] Define the production build handoff: in step 9 a build stage will compile `frontend/` into `dist/`, then the Nginx image will copy those static files. Until then, serve the probe page without requiring a Vite build. Stop routing traffic to the old `frontend:3000` service here; remove that service and its Bun `next start` build/runtime setup during the frontend migration.
- [ ] Define public auth path prefixes explicitly. Adjust authserver-generated redirects, form actions if retained, and cookie paths to work under `/auth`.
- [ ] Add service health endpoints and distinguish liveness from upstream readiness.
- [ ] Strip client-supplied identity/correlation headers at the boundary and overwrite them with gateway-derived values. The backend trusts identity only from the gateway's private service path.

Target routing contract:

| Public route | Destination and behavior |
| --- | --- |
| `GET /dashboard` | Gateway establishes visitor context once step 5 is implemented, then redirects to the canonical `/dashboard/` entry point |
| `GET /dashboard/` and descendants | Gateway proxies the private static Nginx listener; visitor context is added in step 5, and the probe page is replaced by the React/Vite build in step 9; no login required |
| `POST /auth/start` | Middleware starts a browser-bound login transaction |
| `GET /auth/callback` | Middleware validates state, exchanges code, creates app session, redirects |
| `POST /auth/logout` | Middleware clears the app session |
| `POST /auth/reset` | Gateway and authserver cooperate to revoke app/provider sessions and expire their cookies |
| `GET /auth/authorize`, `GET /auth/login` | Authserver demo identity/code flow for the registered transaction |
| `GET /api/observe/snapshot`, `GET /api/observe/stream` | Middleware attaches validated visitor context, then proxies observation backend; implemented in step 8 |
| `GET /api/demo/whoami` | Middleware requires a valid app session, refreshes if necessary, then proxies backend |

Keep token exchange and event ingestion private. Expose public JWKS when it is implemented in step 10. Enforce same-origin/CSRF protection on public state-changing POST routes. Protected API requests should return a useful unauthorized response rather than redirecting a background fetch to HTML; explicit `/auth/start` owns interactive login.

Completion check: the static probe page and protected demo API are reachable through the public gateway without exposing private upstream ports or depending on a frontend JavaScript server. An anonymous protected API request fails; after login it returns the correct gateway-derived demo identity. Observation endpoints become functional in step 8, and the React dashboard replaces the probe page in step 9.

## 5. Establish correlation before adding log displays

- [ ] Have middleware issue an opaque anonymous visitor cookie on the first dashboard request. Store a separate non-secret visitor display ID server-side; the display ID must not authorize access.
- [ ] Generate request IDs at Nginx's public entry point and pass them downstream. Ignore externally supplied request IDs as authoritative IDs. The private static listener preserves that ID rather than generating a second public request identity.
- [ ] Create a new flow ID per login attempt. Associate it with visitor ID, initiating request ID, pending state, client, callback, and eventual provider/app sessions.
- [ ] Register the authorization transaction with authserver over a private authenticated service call so provider events inherit a trusted flow association. Do not accept a browser-selected visitor/flow ID as ownership proof.
- [ ] Emit gateway request events for public dashboard, observation, callback, and protected API requests, linking their Nginx request IDs to the validated visitor context.
- [ ] Join Nginx log entries to these gateway events by request ID. Support either arrival order; access logs arrive after the response is finished.
- [ ] Include the initial dashboard request in that association. It already passed through the gateway that minted the visitor cookie.
- [ ] Determine ownership server-side. Do not identify a visitor by IP address or allow arbitrary visitor/flow IDs in snapshot queries.
- [ ] Keep separate flow IDs for concurrent login attempts, including attempts in different tabs; preserve the correct association through each callback.

Completion check: two browsers on the same network receive distinct visitor contexts, their request/flow associations remain separate, and forged display IDs or correlation headers cannot override ownership. Verify these associations through structured diagnostics now; shared/personal stream isolation is checked after steps 6–8 implement the observation system.

## 6. Implement the observation backend and auth events

Suggested joat packages: `backend/internal/events` for schema/store/subscribers and `backend/internal/observe` for ingestion, snapshots, and stream handlers. Start with bounded in-memory storage; a database is not needed for the first milestone.

- [ ] Define an event envelope with event ID, schema version, timestamp, source service, event type, request ID, visitor/flow references where applicable, outcome, duration, and sanitized details.
- [ ] Assign a monotonic ingestion sequence for stream cursors. Preserve source timestamps and per-flow stage/sequence information; do not rely on clocks alone to order auth stages.
- [ ] Add a private ingestion endpoint authenticated for service producers. Browser routes cannot reach it.
- [ ] Emit structured events at the actual authserver/middleware boundaries: `login_started`, `provider_session_created`/`provider_session_reused`, `code_issued`, `callback_received`, `state_validated`, `code_redeemed`, `tokens_issued`, `app_session_created`, `protected_request_accepted`/`rejected`, refresh, expiry, logout, and failure events.
- [ ] Record initiating flow context before redirecting. Retain intermediate events server-side so navigation cannot lose them.
- [ ] Use bounded buffers and explicit producer timeouts/retries. Observation failures must not indefinitely block login; incomplete traces must show a collection gap rather than invent missing success events.
- [ ] Build separate response projections: sanitized shared traffic and owner-only personal activity/artifact metadata. Do not stream one unrestricted payload and hide fields with frontend filtering.
- [ ] Show token/session fingerprints, token type, expiry, storage location, and cookie names/attributes. Do not include usable session cookies, raw authorization codes, refresh tokens, or secrets in shared logs.
- [ ] If raw disposable access-token inspection is added later, make it an explicit owner-only operation checked against the current app session, not an SSE field broadcast to every viewer.
- [ ] Bound event counts, payload sizes, subscriber queues, and retention. Document that service restarts initially discard in-memory sessions/history.

Completion check: a completed login can be reconstructed from stored events even if no browser had an active stream during the redirects. Personal payloads never enter the shared projection.

## 7. Ingest actual Nginx traffic

Primary files: `nginx/nginx.conf`, `docker-compose.yaml`, and a small collector program in joat, for example `backend/cmd/logcollector`.

- [ ] Extend JSON logs with trusted request ID, normalized path, response status, request duration, and upstream timings where available. Preserve JSON escaping and represent unavailable upstream values safely.
- [ ] Use normalized paths without query strings for the displayed feed; do not log callback codes or tokens into the public projection.
- [ ] Write access logs to an explicit shared volume/file, keeping it separate from container stdout handling. Mount it read-only into the collector.
- [ ] Tail complete JSON lines, cope with malformed lines, truncation, and log rotation, and define a checkpoint/restart policy. Use request IDs to avoid duplicate traffic events after retries/re-reading.
- [ ] Include Nginx-only rejections, 404s, HTTP redirects, and scanner requests, not just requests reaching backend. Configure logging at the relevant HTTP/TLS server scopes.
- [ ] Collect public-entry access logs only. Disable access logging on the private static listener, or keep its diagnostics outside the public feed, so one dashboard request is not counted twice.
- [ ] Mask IP addresses in the shared projection. Treat user agents and paths as untrusted text and limit their lengths.
- [ ] Mark observation polling/streams/assets so the UI can collapse its own background noise. Keep those requests available for inspection without letting them bury scanner and login traffic.

Completion check: the observation store contains sanitized public traffic events for nonexistent paths and Nginx rejections. Ordinary visitor requests join to their owner without exposing visitor-cookie values, and internal static serving does not duplicate traffic. The browser feed is added in steps 8–9.

## 8. Implement snapshot plus persistent SSE

- [ ] Implement `/api/observe/snapshot` and `/api/observe/stream` using standard-library handlers backed by retained events and subscribers. The obsolete single-event placeholder was removed in step 1.
- [ ] Return a snapshot containing recent sanitized traffic, current visitor/session summary, personal flow events, and a stream cursor.
- [ ] Subscribe from that cursor and atomically bridge replay to live delivery so events between snapshot and subscription are not lost or duplicated.
- [ ] Support SSE event IDs and `Last-Event-ID` reconnects; explicitly handle cursors older than retention by requiring a fresh snapshot.
- [ ] Flush event frames, send heartbeats, detach disconnected subscribers, and isolate slow clients with bounded queues.
- [ ] Disable Nginx proxy buffering and caching for SSE and set suitable upstream/read timeouts. Configure the gateway's streaming proxy to flush promptly too.
- [ ] Resolve and authorize visitor context again on reconnect. Stop/reset an existing connection if its personal authorization expires; logout must not leave a private stream authorized forever.
- [ ] After callback redirect, load the completed snapshot before resuming live updates.

Completion check: events arrive incrementally through Nginx, refresh/reconnect resumes correctly, and navigating away during login does not remove the recorded flow. A stalled tab does not stall other viewers.

## 9. Migrate to React/Vite and add the temporary two-column interface

Primary files: replace the current Next.js scaffold with `frontend/index.html`, `frontend/src/main.tsx`, `frontend/src/App.tsx`, `frontend/vite.config.ts`, and associated TypeScript/Tailwind configuration. Update `frontend/package.json`, its lockfile, the frontend build setup, and Compose's static-artifact handoff. No final hero, typography system, animation work, or design overhaul yet.

- [ ] Remove Next.js-specific dependencies, configuration, App Router files, and ESLint configuration. Replace scripts with Vite development/build commands, explicit TypeScript checking, and React-appropriate linting.
- [ ] Retain React, TypeScript, and Tailwind CSS. Set Vite's asset base to `/dashboard/`, and verify asset URLs work after login redirects and direct page reloads.
- [ ] Replace `next/image` and `next/font` usage. Use ordinary image elements/static assets; choose self-hosted fonts during final design rather than depending on Next.js font helpers.
- [ ] Complete the static production build and remove the frontend server container. During development, use Vite's development server with a same-origin proxy arrangement for Go API/auth routes; verify HTTPS/cookie behavior separately through the real Nginx path.
- [ ] Replace the template with two plain log columns: **Server traffic** on the left and **Your activity** on the right. Stack them on narrow screens.
- [ ] Load the snapshot, then connect to SSE. Display timestamps, source, event name, request/flow reference, outcome, and a compact payload preview.
- [ ] Use a small React hook for subscription lifecycle, cleanup, and reconnect/reset state. Start with local component state; add routing or state/data libraries only when concrete UI requirements justify them.
- [ ] Highlight `is_mine` rows in shared traffic using a marker returned by the server. Show all correlated personal request/auth events in the right column, including anonymous activity.
- [ ] Add Login, Make protected request, Logout, and full demo reset controls. Use POST plus same-origin/CSRF protection for state-changing controls.
- [ ] Show anonymous/authenticated state, a safe visitor/session reference, and stream connected/reconnecting status.
- [ ] Expand individual personal rows to inspect sanitized metadata. Render incoming strings as text, never HTML.
- [ ] Bound rendered rows, deduplicate by event ID, and provide pause/follow behavior so new events do not steal the reader's scroll position.
- [ ] Keep durations honest. Show events as they happened; any later replay is explicitly labelled replay rather than artificial live timing.

Completion check: the React/Vite build is served as static files, with no Next.js production runtime. The complete anonymous-to-authenticated experience is usable with these two columns. Login returns to the page and its right column already contains the redirect, code exchange, token, and session sequence.

## 10. Validate the milestone, then add JWT/JWKS depth

First validate the opaque-token milestone locally through Nginx, including a local HTTPS setup for cookie behavior. Repeat the end-to-end checks on the public deployment when deployment is authorized. Deployment is a separate action, not implied by writing this document.

- [ ] Check fresh login, provider-session reuse, protected requests, logout/reset, expired access-token refresh, expired refresh-token rejection, service restart, stream reconnect, and two-browser isolation.
- [ ] Test forged state/ownership headers, code reuse, client/redirect mismatches, missing cookies after another visitor logs in, and unavailable upstreams.
- [ ] Check cookie scope and HTTPS behavior through Nginx, log redaction, collector rotation/restart, retention limits, and event-delivery gaps.
- [ ] Verify the Vite asset base, direct dashboard reloads, navigation fallback, missing-asset 404s, and the gateway-to-static-listener route. Confirm production starts without a frontend Node/Bun/Vite server.
- [ ] Add meaningful automated tests around transaction binding, one-time redemption, refresh rotation, ownership projection, and snapshot/replay handoff; run concurrency checks for shared state.
- [ ] Keep the first milestone labelled OAuth-style demo with opaque tokens. The present middleware checks stored token expiry; it does not perform JWT signature validation.

Then extend the working flow:

- [ ] Issue signed JWT access tokens using an established library. Manage signing keys independently from public JWKS and session storage.
- [ ] Publish JWKS and have the gateway forward the session's server-stored access token to the protected backend. The backend acts as the resource server and validates signature, algorithm, issuer, audience, and expiry before accepting identity. Strip browser-supplied identity/token headers at the gateway; a stored session alone must not bypass backend token validation.
- [ ] Emit separate token-received, signature-validated, claims-validated, and identity-forwarded/backend-accepted events, including actual failures.
- [ ] Add a personal decoded-token view with a clear distinction between decoding and successful validation.
- [ ] Add expired-token, invalid-signature, wrong-audience, and key-rotation demonstrations after the successful JWT path works.
- [ ] Treat OIDC as a separate extension: implement ID tokens, nonce handling, discovery, and relevant validation before presenting the system as an OIDC provider/client.

Completion check: the logs distinguish token issuance, decoding, validation, session lookup, refresh, and authorization decisions accurately. Failed validation prevents protected access.

## 11. Build the final dashboard and design last

Use `docs/design.md` and `docs/style.md` once the two-column interface proves the event model and behavior.

- [ ] Update the architecture documentation to replace oauth2-proxy with notauth's actual components and implemented protocol capabilities.
- [ ] Update stack descriptions to show a `net/http` backend and React/Vite static frontend. Apply the visual direction independently of Next.js; choose reusable React components only where they help the actual inspectors and controls.
- [ ] Turn the personal trace into the hero: anonymous connection before login; complete, selectable auth timeline after login.
- [ ] Keep shared Nginx traffic visible as the surrounding live context and preserve the ownership highlight.
- [ ] Add linked session, token, cookie, forwarded-header, and request inspectors using the existing owner-scoped data contracts.
- [ ] Add the glowing Login button, restrained dark instrument styling, responsive layout, accessible focus/state labels, and reduced-motion behavior.
- [ ] Add explicit replay and selected failure experiments without changing the underlying evidence into scripted success animations.

## First milestone acceptance criteria

- [ ] The backend uses standard-library HTTP handlers without Gin, and the frontend is a React/Vite static build without a Next.js runtime.
- [ ] The public page works without login and both columns update continuously.
- [ ] Real Nginx traffic appears, including requests that never reach the application.
- [ ] The visitor's initial anonymous request and subsequent activity are identifiable.
- [ ] Clicking Login requires no input and returns to the same dashboard.
- [ ] The personal column contains the complete actual login trace despite browser navigation.
- [ ] The protected demo endpoint returns the correct identity and rejects anonymous access.
- [ ] Tokens, sessions, cookies, and refresh behavior are inspectable through accurate, safely scoped metadata.
- [ ] Two simultaneous visitors cannot read each other's private traces or credentials.
- [ ] Missing events, failures, restarts, and stream disconnections are visible rather than disguised.
- [ ] Final frontend design has not delayed any of the above.
