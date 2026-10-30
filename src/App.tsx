import { KageLandingPage } from "@designcodeio/threeui";
import "@designcodeio/threeui/style.css";

export function Scene() {
  return (
    <div className="shader-frame" style={{ 
      padding: '0',
      width: '100%',
      height: '100%',
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center'
    }}>
      <KageLandingPage
        headingFont="onest"
        bodyFont="onest"
        headingWeight="400"
        bodyWeight="300"
        primaryColor="#e0231c"
        headingSize={46}
        bodySize={17}
        headingLetterSpacing={-0.012}
        style={{
          width: '95%',
          height: '95%'
        }}
      />
    </div>
  );
}

function App() {
  return (
    <div style={{ 
      width: "100vw", 
      height: "100vh",
      overflow: 'hidden'
    }}>
      <Scene />
    </div>
  );
}

export default App;