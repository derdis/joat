# Design Guide

## Purpose

This document turns the style analysis in [style.md](/home/alex/code/joat/docs/style.md) into a project-specific design direction for Jack of All Trades Lab.

The product is not a generic developer dashboard. It is a live infrastructure and IAM lab whose UI should make authentication and request flow visible, inspectable, and easy to follow in real time.

The design goal is:

> Present the system like a serious black-box operating surface, then reveal auth and network behavior as live, precise instrumentation.

## Product Context

Jack of All Trades Lab exists to show:

- reverse proxy flow
- OIDC authentication
- JWT validation
- identity propagation
- request lifecycle visibility
- real-time system events

The frontend should make those invisible steps legible. Every major UI decision should support understanding, not decoration.

## Design Thesis

Use the conclusions from `docs/style.md` directly:

- Factory defines the shell, typography hierarchy, spacing, chrome, and overall authority.
- Databuddy defines the motion model and the feel of live product surfaces inside the dashboard.

In practice:

- the overall frame is stark, monochrome, sparse, and terminal-adjacent
- the product modules are dense, reactive, and quietly animated
- orange is used as a precise signal color, not as broad branding
- content must feel operational, not promotional

## Experience Model

The product should be designed as two connected layers.

### 1. Outer Layer: System Shell

This is the landing and framing layer.

It should communicate:

- authority
- technical credibility
- infrastructure focus
- clarity of purpose

Visual behavior:

- black or near-black stage
- large left-weighted hero
- compact utility navigation
- terminal-style CTA or quickstart block
- restrained trust or architecture references
- large areas of empty space

### 2. Inner Layer: Live Instrument Panel

This is the core product surface.

It should communicate:

- request flow in motion
- authentication state transitions
- token and header visibility
- timing and sequencing
- observability of infrastructure behavior

Visual behavior:

- dark framed modules
- thin borders
- compact labels
- monospaced metadata
- smooth event updates
- subtle glow on active lines and states

## Primary Use Cases

The design should support these user tasks first:

1. Understand where a request is in the auth flow.
2. See whether the request is anonymous, redirected, authenticated, or failed.
3. Inspect decoded JWT data and injected identity headers.
4. Review request metadata such as IP, latency, and timestamps.
5. Watch the event stream update live without losing readability.
6. Understand the system architecture from the landing page before entering the dashboard.

## Information Architecture

The project should be structured as a focused single-product experience.

### Landing Page

Recommended order:

1. thin utility nav
2. hero with explicit promise
3. terminal quickstart or status block
4. low-key architecture or trust strip
5. product preview section
6. feature walkthrough
7. architecture explanation
8. closing CTA
9. utility footer

### Dashboard / Lab View

Recommended layout:

1. top system bar
2. live status strip
3. primary timeline area
4. side inspection panels
5. lower detail modules

Core modules:

- request stream
- active request timeline
- auth step log
- JWT viewer
- identity header panel
- request metadata panel
- system status / services panel

## Landing Page Direction

### Hero

The hero should follow the Factory-led composition from `docs/style.md`.

Recommended content model:

- eyebrow: small, utilitarian, system-like
- headline: direct and literal
- support copy: mono, short, operational
- CTA block: terminal or command-framed
- ambient artifacts: small floating cards, request samples, or auth-state indicators

Good headline territory:

- `Watch authentication happen in real time.`
- `Visible infrastructure for identity and request flow.`
- `A live lab for reverse proxy, OIDC, and JWT behavior.`

Avoid:

- abstract AI claims
- playful metaphors
- broad platform language
- startup-brand optimism

### Architecture Section

This section should explain the stack clearly:

`Internet -> Nginx -> oauth2-proxy -> Backend -> Event stream -> Frontend`

Presentation guidance:

- horizontal flow on desktop
- stacked flow on mobile
- framed node blocks
- thin connector lines
- optional active-state pulse on live systems

### Product Preview

This section should preview the lab as an instrument panel, not as a marketing card grid.

Show:

- a request timeline
- an event feed
- a decoded token block
- identity headers

The preview should feel executable and operational.

## Dashboard Direction

### Overall Layout

The dashboard should feel denser than the landing page but still controlled.

Recommended desktop structure:

- left/main column for request timeline and event stream
- right column for JWT, headers, and request inspection
- lower row for service health, counters, and architecture state

Recommended mobile structure:

- stacked modules
- timeline first
- inspector second
- metadata and service panels last

### Top Bar

Purpose:

- establish environment
- expose current system state
- anchor key actions

Contents:

- project name
- environment label
- connection status
- auth provider status
- backend status
- clear entry to replay or inspect recent flows later

### Live Status Strip

This strip should surface the current operational state at a glance.

Suggested signals:

- active requests
- authenticated requests
- redirects in progress
- recent failures
- event stream health

These numbers should animate smoothly, not jump harshly.

### Request Timeline

This is the central module of the product.

Each request should read like a traceable sequence:

- request received
- auth required
- redirect issued
- login completed
- token received
- token validated
- headers injected
- backend served response

Design rules:

- each request has a stable row or lane
- steps appear progressively
- timestamps are monospaced
- active step is highlighted with subtle glow
- success, redirect, and failure states are visually distinct without becoming colorful noise

### Event Feed

The event feed should feel continuous and reliable.

Rules:

- newest events visible immediately
- no aggressive auto-scroll when the user is reading older content
- monospaced rows
- compact severity markers
- event type is the primary label
- payload preview is secondary and truncated

Orange should only mark active or attention states. Failures can use muted red sparingly, but the overall palette should remain monochrome-first.

### JWT Viewer

The JWT panel should feel inspectable and educational.

Display:

- raw token presence
- decoded header
- decoded payload
- expiration
- issuer
- subject
- email or identity claims when available

Design rules:

- code-style formatting
- clear segmentation between header, payload, and metadata
- scrollable but contained
- sensitive feeling, high-precision presentation

### Identity Header Panel

This module should show exactly what is forwarded downstream.

Include:

- header name
- header value
- source stage if known

The framing should make propagation explicit. The user should immediately understand what the proxy added before the backend handled the request.

### Request Metadata Panel

Show:

- request ID
- path
- method
- IP
- response status
- upstream latency
- total duration

This module should be compact and easy to scan, using mono for values.

### Service Health / Architecture Panel

This panel explains whether the moving pieces are healthy.

Suggested services:

- nginx
- oauth2-proxy
- backend
- event stream

Presentation:

- simple bordered rows
- tiny status dots
- last update timestamp
- optional low-key sparkline or recent activity trace

## Visual System

### Palette

Use the token direction from `docs/style.md` as the base:

- background: `#050505`
- surface: `#101010`
- elevated surface: `#151515`
- primary text: `#f5f5f0`
- secondary text: `#8d8d87`
- border: `rgba(255,255,255,0.1)`
- accent: `#f47a20`

Extended state colors should stay restrained:

- success: muted off-white or very soft green tint
- warning: accent orange
- error: muted dark red, used sparingly
- live glow: subtle white bloom, not neon color

### Typography

Use two-mode typography:

- sans serif for headlines, navigation, and structural UI
- monospace for supporting copy, timestamps, payloads, labels, and system output

Rules:

- headlines should be oversized and tightly controlled
- section headings should be compact and crisp
- body copy should be short and declarative
- metadata should feel like interface output, not editorial prose

### Spacing

The layout should alternate between spacious and dense.

Rules:

- large vertical breathing room in hero and architecture sections
- tight, instrument-like spacing inside dashboard panels
- avoid equally padded sections everywhere
- let empty space create authority around the product modules

### Borders and Radius

Use:

- thin visible borders
- restrained radius
- low shadow
- quiet separators

Avoid:

- soft oversized cards
- glassmorphism
- heavy blur
- glossy surfaces

## Motion

Motion should follow the Databuddy influence inside product modules and stay subtle everywhere else.

Use:

- faint fades
- prompt blink or cursor pulse
- smooth counter updates
- gentle chart interpolation
- subtle active-line glow
- quiet module reveal on load

Avoid:

- spring-heavy motion
- bounce
- loud parallax
- jittery real-time updates
- constant decorative movement

Recommended timings:

- fast: `140ms ease`
- base: `220ms ease`
- slow: `420ms ease`

## Interaction Rules

### Live Updating

The UI must make real-time behavior legible.

Rules:

- preserve row stability when new events arrive
- animate inserts softly
- do not reorder visible items unexpectedly
- allow pausing or reviewing historical entries later
- never let live updates make payload inspection difficult

### Hover and Focus

Interactions should feel exact and low-noise.

Use:

- border brightening
- subtle background lift
- tiny indicator changes

Avoid:

- oversized lift
- glow explosions
- playful button behavior

### Empty and Error States

These states should still feel like system output.

Examples:

- `Waiting for first authenticated request.`
- `No JWT available for this request.`
- `Event stream disconnected. Retrying.`

Avoid cute illustrations or friendly placeholder copy.

## Copy Direction

Copy should be:

- declarative
- technical
- compressed
- credible
- low-hype

Preferred vocabulary:

- request
- redirect
- provider
- token
- validate
- propagate
- identity
- proxy
- upstream
- event stream
- latency
- session

Avoid:

- magical language
- vague productivity claims
- broad platform slogans
- consumer-friendly cheerfulness

## Responsive Behavior

### Desktop

Desktop should preserve the left-weighted hero and asymmetry from the style guide.

The dashboard should use multi-column layouts where the timeline remains dominant and inspectors remain secondary.

### Mobile

Mobile should preserve seriousness, not collapse into generic stacked cards with oversized spacing.

Rules:

- keep typography bold but controlled
- stack modules in order of diagnostic value
- keep metadata dense
- retain monospaced labels where useful
- simplify ambient artifacts instead of shrinking everything indiscriminately

## Accessibility

The interface is dark and high-contrast, so accessibility needs explicit handling.

Requirements:

- maintain strong contrast for all text and borders
- do not rely on color alone for state
- preserve keyboard focus visibility
- respect reduced motion preferences
- ensure timestamps and payload text remain readable at small sizes

## Current Repo Implication

The current frontend is still scaffold-level and does not reflect the intended design system yet. The next implementation phase should replace the template UI with:

1. a Factory-led landing shell
2. a live auth-flow dashboard surface
3. a shared token system for color, type, spacing, and motion

The design work should prioritize product comprehension over polish theater.

## Final Direction

If this document is followed correctly, Jack of All Trades Lab should feel like:

- a severe developer-facing operating surface on the outside
- a live observability instrument on the inside
- a product that teaches authentication and infrastructure behavior by making every state visible

In one sentence:

Build Jack of All Trades Lab like a black, high-authority infrastructure console whose live panels make auth, proxy, and identity flow understandable and observable at a glance.
