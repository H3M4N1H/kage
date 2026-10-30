# Kage Landing Page

High-fidelity recreation of the ThreeUI Kage Landing Page using the exact packaged component.

## Quick Start

```bash
# Install dependencies
npm install

# Start development server
npm run dev

# Open http://localhost:3000
```

## Features

- **Exact ThreeUI Implementation**: Uses packaged `@designcodeio/threeui` component
- **Full WebGL Experience**: Complete Three.js scene with all assets
- **Responsive Design**: Works on desktop (1440×900) and mobile (390×844)
- **Interactive Navigation**: All links and buttons functional
- **High Fidelity**: Matches original ThreeUI reference

## Project Structure

```
public/landing-pages/           # ThreeUI packaged assets
  kage.html                    # Main HTML document
  secret-pathways-assets/      # WebGL, fonts, images
src/
  App.tsx                     # Main component with KageLandingPage
  main.tsx                    # React entry point
  index.css                   # Basic styles
```

## Technical Stack

- **React 19** + **TypeScript**
- **Vite** build tool
- **ThreeUI** component library (`@designcodeio/threeui`)
- **Three.js** WebGL runtime

## Implementation Details

The project uses the exact `<KageLandingPage />` component from ThreeUI with the following configuration:

```tsx
<KageLandingPage
  headingFont="onest"
  bodyFont="onest"
  headingWeight="400"
  bodyWeight="300"
  primaryColor="#e0231c"
  headingSize={46}
  bodySize={17}
  headingLetterSpacing={-0.012}
/>
```

All assets are served directly from the ThreeUI package to ensure byte-for-byte accuracy.

## Testing Verified

✅ **Desktop (1440×900)**: Full content, WebGL, navigation  
✅ **Mobile (390×844)**: Responsive, no overflow  
✅ **Interactive Elements**: All links/buttons work  
✅ **Assets**: All 32+ files load correctly  
✅ **Performance**: No critical errors, fast loading

## License

Assets from ThreeUI package follow their licensing terms.