// Dependency-free, perspective-projected 3D point cloud. No visitor data leaves the page.
class PortfolioGlobe extends HTMLElement {
  async connectedCallback() {
    this.abort = new AbortController();
    const {signal} = this.abort;
    this.canvas = this.querySelector('canvas');
    this.ctx = this.canvas.getContext('2d');
    if (!this.ctx) return;
    this.longitude = 80.27;
    this.latitude = 16;
    this.zoom = 1;
    this.spin = false;
    this.visible = true;
    this.motion = matchMedia('(prefers-reduced-motion: reduce)');
    try {
      const response = await fetch('globe-points.json', {signal});
      if (!response.ok) throw new Error('Globe data unavailable');
      const coordinates = await response.json();
      this.points = coordinates.map(([lon, lat]) => {
        const a = lon * Math.PI / 180, b = lat * Math.PI / 180;
        return {x: Math.cos(b) * Math.sin(a), y: Math.sin(b), z: Math.cos(b) * Math.cos(a), accent: lon > 65 && lon < 92 && lat > 6 && lat < 34};
      });
      this.surface = [];
      for (let i=0;i<7000;i++) {
        const y=1-2*(i+.5)/7000, a=i*Math.PI*(3-Math.sqrt(5)), r=Math.sqrt(1-y*y);
        this.surface.push({x:r*Math.cos(a),y,z:r*Math.sin(a)});
      }
      if (signal.aborted) return;
      this.canvas.hidden = false;
      this.querySelector('.globe-fallback').hidden = true;
      this.dataset.ready = 'true';
      this.resizeObserver = new ResizeObserver(() => this.draw());
      this.resizeObserver.observe(this.querySelector('.globe-stage'));
      this.visibilityObserver = new IntersectionObserver(([entry]) => {
        this.visible = entry.isIntersecting;
        this.animate();
      });
      this.visibilityObserver.observe(this);
      document.addEventListener('visibilitychange', () => this.animate(), {signal});
      this.motion.addEventListener('change', () => {if(this.motion.matches) this.setSpin(false);}, {signal});
      this.canvas.addEventListener('keydown', event => {
        const actions = {' ':'spin', Enter:'spin', ArrowLeft:'left', ArrowRight:'right', ArrowUp:'up', ArrowDown:'down', '+':'zoom-in', '=':'zoom-in', '-':'zoom-out', Home:'reset'};
        if (actions[event.key]) {event.preventDefault();this.action(actions[event.key]);}
      }, {signal});
      this.canvas.addEventListener('pointerdown', event => {
        if (this.drag || (event.pointerType === 'mouse' && event.button !== 0)) return;
        const wasSpinning = this.spin;
        this.setSpin(false);
        this.drag = {id:event.pointerId, x:event.clientX, y:event.clientY, startX:event.clientX, startY:event.clientY, moved:false, wasSpinning};
        this.canvas.setPointerCapture(event.pointerId);
        this.canvas.classList.add('dragging');
      }, {signal});
      this.canvas.addEventListener('pointermove', event => {
        if (!this.drag || this.drag.id !== event.pointerId) return;
        const drag = this.drag;
        if (!drag.moved && Math.hypot(event.clientX-drag.startX,event.clientY-drag.startY)<4) return;
        drag.moved = true;
        this.longitude -= (event.clientX - drag.x) * .35;
        this.latitude = Math.max(-70, Math.min(70, this.latitude + (event.clientY - drag.y) * .3));
        drag.x = event.clientX;drag.y = event.clientY;
        this.draw();
      }, {signal});
      const endDrag = event => {
        if (!this.drag || this.drag.id !== event.pointerId) return;
        const drag = this.drag;
        this.drag = null;
        this.canvas.classList.remove('dragging');
        if (this.canvas.hasPointerCapture(event.pointerId)) this.canvas.releasePointerCapture(event.pointerId);
        if (event.type === 'pointerup' && !drag.moved) this.setSpin(!drag.wasSpinning);
      };
      for (const name of ['pointerup','pointercancel','lostpointercapture']) this.canvas.addEventListener(name, endDrag, {signal});
      this.draw();
      this.setSpin(!this.motion.matches);
    } catch (error) {
      if (error.name !== 'AbortError') this.querySelector('#globe-help').textContent = 'Chennai, India · Static globe preview';
    }
  }
  disconnectedCallback() {
    this.abort?.abort();this.resizeObserver?.disconnect();this.visibilityObserver?.disconnect();cancelAnimationFrame(this.frame);
  }
  setSpin(spin) {
    this.spin = spin;
    this.dataset.spinning = String(spin);
    this.canvas.setAttribute('aria-pressed', String(spin));
    this.canvas.setAttribute('aria-label', spin ? 'Pause globe rotation' : 'Resume globe rotation');
    this.animate();
  }
  action(action) {
    if (action === 'spin') {this.setSpin(!this.spin);return;}
    this.setSpin(false);
    if (action === 'left') this.longitude -= 15;
    if (action === 'right') this.longitude += 15;
    if (action === 'up') this.latitude = Math.min(70, this.latitude + 10);
    if (action === 'down') this.latitude = Math.max(-70, this.latitude - 10);
    if (action === 'zoom-in') this.zoom = Math.min(1.3, this.zoom + .1);
    if (action === 'zoom-out') this.zoom = Math.max(.7, this.zoom - .1);
    if (action === 'reset') {this.longitude = 80.27;this.latitude = 16;this.zoom = 1;}
    this.draw();
  }
  animate() {
    cancelAnimationFrame(this.frame);
    if (!this.spin || !this.visible || document.hidden) return;
    this.lastTime = performance.now();
    const tick = now => {
      this.longitude += Math.min(now - this.lastTime, 60) * .004;
      this.lastTime = now;this.draw();
      this.frame = requestAnimationFrame(tick);
    };
    this.frame = requestAnimationFrame(tick);
  }
  draw() {
    if (!this.points) return;
    const rect = this.querySelector('.globe-stage').getBoundingClientRect();
    const w = rect.width, h = rect.height;
    if (!w || !h) return;
    const dpr = Math.min(devicePixelRatio || 1, 2);
    if (this.canvas.width !== Math.round(w*dpr) || this.canvas.height !== Math.round(h*dpr)) {
      this.canvas.width = Math.round(w*dpr);this.canvas.height = Math.round(h*dpr);
    }
    const ctx = this.ctx;ctx.setTransform(dpr,0,0,dpr,0,0);ctx.clearRect(0,0,w,h);
    const radius = Math.min(w,h)*.40*this.zoom;
    const yaw = this.longitude*Math.PI/180, tilt = this.latitude*Math.PI/180;
    const sy=Math.sin(yaw),cy=Math.cos(yaw),st=Math.sin(tilt),ct=Math.cos(tilt);
    const project = p => {
      const x=p.x*cy-p.z*sy,z=p.x*sy+p.z*cy;
      const y=p.y*ct-z*st,depth=p.y*st+z*ct;
      const perspective=4/(4-depth);
      return {x:w/2+x*radius*perspective,y:h/2-y*radius*perspective,z:depth,scale:perspective};
    };
    const spherePoint = (lon,lat) => ({x:Math.cos(lat)*Math.sin(lon),y:Math.sin(lat),z:Math.cos(lat)*Math.cos(lon)});
    // Orbit planes stay anchored to the camera while the geographic sphere rotates.
    const orbitPoint = (t, angle, inclination, size) => {
      const x=Math.cos(t)*size,y=Math.sin(t)*Math.sin(inclination)*size,z=Math.sin(t)*Math.cos(inclination)*size;
      const scale=4/(4-z);
      return {x:w/2+(x*Math.cos(angle)-y*Math.sin(angle))*radius*scale,
        y:h/2+(x*Math.sin(angle)+y*Math.cos(angle))*radius*scale,z};
    };
    const orbits=[[-.92,.24,1.2],[.53,.33,1.16],[-.15,.11,1.12]];
    const drawOrbits = front => {
      for (let j=0;j<orbits.length;j++) {
        const [angle,inclination,size]=orbits[j];
        ctx.beginPath();let started=false;
        for(let n=0;n<=240;n++) {
          const p=orbitPoint(n*Math.PI/120,angle,inclination,size);
          if((p.z>=0)!==front){started=false;continue;}
          if(!started){ctx.moveTo(p.x,p.y);started=true;}else ctx.lineTo(p.x,p.y);
        }
        ctx.strokeStyle=front ? ['#b6ff1880','#b6ff184d','#b6ff1838'][j] : '#b6ff182b';
        ctx.lineWidth=j===0?.8:.55;
        ctx.setLineDash(j===2?[2,7]:[]);ctx.stroke();ctx.setLineDash([]);
        for(const t of [.6,2.8,4.9]){
          const p=orbitPoint(t+j*.7,angle,inclination,size);
          if((p.z>=0)!==front)continue;
          ctx.fillStyle=front?'#c1ff00':'#526d21';ctx.fillRect(p.x-1,p.y-1,2,2);
        }
      }
    };
    drawOrbits(false);
    const edge=radius*4/Math.sqrt(15);
    const shade=ctx.createRadialGradient(w/2-radius*.3,h/2-radius*.3,0,w/2,h/2,edge);
    shade.addColorStop(0,'#10180d');shade.addColorStop(.65,'#090e0b');shade.addColorStop(1,'#050807');
    ctx.beginPath();ctx.arc(w/2,h/2,edge,0,Math.PI*2);ctx.fillStyle=shade;ctx.fill();
    ctx.strokeStyle='#63715366';ctx.lineWidth=.6;ctx.stroke();
    // Dim ocean dots preserve the sphere's volume on the Pacific-facing side.
    ctx.fillStyle='#7e9473';
    for(const point of this.surface){
      const p=project(point);if(p.z<.02)continue;
      ctx.globalAlpha=.08+.17*p.z;
      ctx.fillRect(p.x,p.y,.65*p.scale,.65*p.scale);
    }
    ctx.globalAlpha=1;ctx.lineWidth=.45;
    const grid = (latitude,fixed) => {
      ctx.beginPath();let started=false;
      for(let n=-180;n<=180;n+=3){
        const p=project(spherePoint((latitude?n:fixed)*Math.PI/180,(latitude?fixed:n/2)*Math.PI/180));
        if(p.z<0){started=false;continue;}
        if(!started){ctx.moveTo(p.x,p.y);started=true;}else ctx.lineTo(p.x,p.y);
      }
      ctx.strokeStyle='#53654738';ctx.stroke();
    };
    for(let lat=-60;lat<=60;lat+=30)grid(true,lat);
    for(let lon=-180;lon<180;lon+=30)grid(false,lon);
    for(const point of this.points){
      const p=project(point);if(p.z<.02)continue;
      ctx.globalAlpha=.25+.65*p.z;
      ctx.fillStyle=point.accent?'#c1ff00':'#cdd4cf';
      ctx.beginPath();ctx.arc(p.x,p.y,(.35+.25*p.z)*p.scale,0,Math.PI*2);ctx.fill();
    }
    ctx.globalAlpha=1;
    drawOrbits(true);
    const chennai=project(spherePoint(80.27*Math.PI/180,13.08*Math.PI/180));
    if(chennai.z>0){
      ctx.strokeStyle='#c1ff00';ctx.fillStyle='#c1ff00';ctx.lineWidth=1;
      ctx.beginPath();ctx.arc(chennai.x,chennai.y,7,0,Math.PI*2);ctx.stroke();
      ctx.fillRect(chennai.x-2,chennai.y-2,4,4);
      const x=Math.min(w-102,chennai.x+20),y=chennai.y+28;
      ctx.beginPath();ctx.moveTo(chennai.x+7,chennai.y+7);ctx.lineTo(x,y-5);ctx.stroke();
      ctx.fillStyle='#050505';ctx.fillRect(x-4,y-3,100,22);
      ctx.fillStyle='#c1ff00';ctx.font='11px JetBrains,monospace';ctx.fillText('CHENNAI, IN',x,y+12);
    }
    this.dataset.longitude = this.longitude.toFixed(2);
    this.dataset.zoom = this.zoom.toFixed(1);
  }
}
customElements.define('portfolio-globe',PortfolioGlobe);
