// Set the reveal state before the first paint; fail open if the main script fails.
if (!matchMedia('(prefers-reduced-motion: reduce)').matches && 'IntersectionObserver' in window) {
  document.documentElement.classList.add('motion-ready');
  window.motionFallback = setTimeout(() => document.documentElement.classList.remove('motion-ready'), 4000);
}
