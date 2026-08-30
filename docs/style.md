# Style Guide

## Basis

This guide is based on:

- the live homepages of `databuddy.cc` and `factory.ai`
- the provided screenshots, with emphasis on Factory's visual design

The key adjustment after reviewing the screenshots is this:

- Factory defines the brand style
- Databuddy contributes the product-motion reference: real-time UI, soft glow, smooth frontend behavior

This should not be treated as a 50/50 blend. If there is a conflict, follow Factory for brand, layout, typography, and chrome. Use Databuddy mainly for interaction feel inside product surfaces.

## Core Direction

The target style is:

- stark
- technical
- monochrome-first
- high-contrast
- terminal-adjacent
- precise
- spacious in hero layout
- dense inside product modules

It should feel like software infrastructure for serious builders, not a friendly startup dashboard.

## Primary Reference: Factory

Factory's design is the dominant reference system.

### Brand Character

Factory feels:

- severe
- controlled
- intelligent
- operational
- anti-decorative

The page does not try to charm. It tries to establish authority.

### Visual Signature

The standout traits in Factory's design are:

- nearly black background
- oversized white headlines
- very sparse use of color
- orange as a precise signal color, not a broad brand wash
- lots of negative space
- thin UI outlines
- terminal-window framing
- monospaced supporting text
- tiny utility navigation
- small floating interface artifacts in the background

This is not generic dark SaaS. It is closer to a stripped-down operating environment.

### Layout Rules

Factory uses asymmetry and emptiness deliberately.

- Keep the hero left-weighted.
- Let large parts of the canvas stay empty.
- Use floating micro-elements on the right as ambient technical artifacts.
- Avoid filling every area with cards or illustrations.
- Let one major idea dominate each screen.

The visual effect should be confidence through restraint.

### Hero Composition

Factory's hero establishes the system:

- tiny eyebrow label
- very large headline
- monospaced explanatory copy
- terminal-like command block
- sparse ambient UI fragments
- subdued trust logos near the fold

This composition should be preserved.

Recommended hero structure:

1. small uppercase or utility eyebrow
2. oversized 2 to 4 line headline
3. short monospace paragraph
4. terminal CTA module
5. low-contrast logo strip

### Typography

Factory's typography is the most important stylistic clue.

Use two text modes:

- a clean sans serif for headlines and navigation
- a monospace face for supporting copy, terminal UI, and system labels

Rules:

- Headlines should be very large, smooth, and tightly controlled.
- Supporting paragraphs can switch to monospace for a more computational tone.
- Navigation should be compact and understated.
- Small labels should feel like interface metadata, not editorial category tags.

Typography should create a conversation between product marketing and terminal UX.

### Color System

Factory's palette should be interpreted as:

- `#000000` to near-black for background
- white or warm-white for primary text
- muted gray for secondary content
- a single orange accent for active state, bullets, prompts, and tiny indicators

Orange is not used for broad gradients or backgrounds. It is used like a cursor or LED.

Use orange for:

- prompt symbols
- active dots
- small status markers
- selective emphasis in tiny components

Do not use orange for:

- large fills
- full buttons everywhere
- heavy backgrounds
- decorative glows

### Navigation

Factory's nav is compact, uppercase-leaning, and utility-first.

Guidelines:

- keep it thin
- keep the type small
- spread links horizontally
- use simple bordered buttons on the right
- avoid oversized nav branding

The nav should feel like software control chrome.

### UI Components

Factory's UI surfaces are not soft app cards. They are framed instruments.

Use:

- thin borders
- dark panels
- restrained radius
- low shadow
- quiet separators
- command-line framing

Preferred modules:

- terminal install box
- tabbed OS selector
- prompt field
- compact command snippets
- small floating code/task cards
- simple grid indicators

Avoid:

- soft glassmorphism
- colorful charts as the main visual identity
- oversized pill-heavy design systems
- playful illustrations

### Motion

Factory motion should feel subtle and machine-like.

Use:

- faint fades
- cursor or prompt blink
- quiet hover states
- tab transitions
- slight reveal of floating artifacts

Avoid:

- springy motion
- bouncy cards
- dramatic parallax
- loud interaction feedback

## Secondary Reference: Databuddy

Databuddy should mainly influence how live product surfaces behave.

### What Actually Stands Out

The key differentiator in the Databuddy screenshot is not the outer shell. It is the feeling of:

- constant live updates
- soft real-time glow
- fluid chart motion
- polished dashboard transitions
- smoothness across dense UI modules

The interface feels alive without feeling noisy.

### Product Surface Traits

Use Databuddy as the reference for:

- real-time graphs
- updating metric cards
- live session lists
- funnels
- error panels
- feature flag rows
- event counters

The modules are dense, but they remain calm because motion is continuous and soft rather than flashy.

### Visual Characteristics

Databuddy's product surface suggests:

- dark graphite panel backgrounds
- subtle inner glow or bloom around active lines
- light gray data visualization on dark canvas
- thin borders
- compact spacing
- highly legible module titles

The line chart in particular matters:

- glowing white line
- smooth curve
- understated area fade beneath
- no loud multicolor treatment

### Interaction Feel

If a page or app borrows from Databuddy, the frontend should feel:

- fast
- smooth
- low-jank
- quietly animated
- continuously updating

Use:

- eased transitions
- live number refresh
- smooth chart interpolation
- subtle shimmer or pulse for active states
- scrolling that feels stable and polished

Avoid:

- abrupt metric jumps
- noisy flashing indicators
- excessive skeleton loading
- gimmicky dashboard animation

## Combined Style Definition

When combining the two references, the intended result is:

- Factory for shell
- Databuddy for live product motion

In practice that means:

- Use Factory's black-stage presentation for the page frame.
- Use Factory's oversized headline and mono-support text model.
- Use Factory's orange micro-accent system.
- Use Databuddy's smooth, glowing, real-time behavior inside product demos.
- Keep all product modules restrained and precise.

The exterior should feel severe. The interior should feel alive.

## Layout Blueprint

### Overall Page

1. Thin utility nav on a black background
2. Large left-aligned hero with oversized headline
3. Terminal-style CTA/install block
4. Low-key logo strip
5. Product sections with dark framed modules
6. Dense feature walkthrough
7. Enterprise/trust section
8. Strong but minimal closing CTA
9. Utility footer

### Section Rhythm

Alternate between:

- spacious conceptual sections
- denser product sections

Do not make every section equally dense. Factory works because the emptiness gives the product modules weight.

## Component Rules

### Buttons

Buttons should be simple and sharp.

- white or transparent
- small to medium height
- restrained radius
- compact labels
- clear border or solid contrast

Preferred button language:

- `Log In`
- `Contact Sales`
- `Start Free`
- `Quickstart`
- `Learn More`

### Eyebrows and Labels

Use small labels like system markers.

- uppercase or near-uppercase
- mono or compact sans
- high letter spacing when appropriate
- tiny accent bullet optional

They should feel like interface metadata.

### Cards and Panels

Cards should feel instrument-like.

- dark surface
- visible border
- low decoration
- compact internal spacing
- one clear purpose per panel

### Demo Blocks

Demo blocks should look executable.

Good patterns:

- command entry
- prompt line
- active tab state
- code task snippets
- small floating annotation boxes
- real-time dashboard fragments

## Copy Style

### Voice

The copy should be:

- declarative
- technical
- compressed
- credible
- low-hype

### Headline Rules

- Make the category obvious.
- Keep the promise direct.
- Avoid metaphor-heavy startup phrasing.
- Prefer strong nouns over abstract inspiration.

Good examples of the style:

- `Agent-Native Software Development`
- `Real-time analytics that runs itself`
- `Droids where you code`
- `Live product intelligence without the overhead`

### Supporting Text

Supporting copy can be slightly more mechanical than typical marketing text.

- monospace is acceptable
- short paragraphs
- operational outcomes
- explicit workflow references

Mention:

- IDE
- CI/CD
- sessions
- vitals
- events
- refactors
- migrations
- flags

That vocabulary is part of the style.

## Design Tokens

### Factory-Led Tokens

- `--bg: #050505`
- `--surface: #101010`
- `--surface-2: #151515`
- `--text: #f5f5f0`
- `--muted: #8d8d87`
- `--line: rgba(255,255,255,0.1)`
- `--accent: #f47a20`
- `--radius-sm: 8px`
- `--radius-md: 12px`
- `--radius-lg: 16px`
- `--shadow: 0 8px 24px rgba(0,0,0,0.28)`
- `--container: 1280px`

### Databuddy Motion Tokens

- `--live-glow: 0 0 18px rgba(255,255,255,0.22)`
- `--chart-line: #f2f2f2`
- `--chart-fill: linear-gradient(to bottom, rgba(255,255,255,0.18), rgba(255,255,255,0.02))`
- `--transition-fast: 140ms ease`
- `--transition-base: 220ms ease`
- `--transition-slow: 420ms ease`

## Do and Don't

### Do

- build the page around a black canvas
- use very large white headlines
- use monospace strategically
- keep the accent color scarce and intentional
- make product demos feel executable
- use negative space confidently
- make live data feel smooth and ambient
- keep panels dense and clean

### Don't

- fill the page with colorful gradients
- use playful AI illustration tropes
- over-round components
- make every section equally busy
- add soft consumer-brand friendliness
- use loud dashboard rainbow charts
- animate metrics in a jittery way

## Final Definition

The style should be read as:

Factory's stark terminal-modern brand system

with

Databuddy's smooth real-time frontend behavior inside the product layer

If you need a single sentence:

Build the page like a black, high-authority developer operating system, then make the product surfaces feel quietly live.
