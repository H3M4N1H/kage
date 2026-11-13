const filters = document.querySelectorAll('[data-filter]');
const projects = document.querySelectorAll('[data-category]');
const moreProjects = document.querySelector('.more-projects');
function filterProjects(button) {
  moreProjects.hidden = true;
  for (const item of filters) {
    const selected = item === button;
    item.classList.toggle('active', selected);
    item.setAttribute('aria-pressed', String(selected));
  }
  let count = 0;
  for (const project of projects) {
    project.hidden = button.dataset.filter !== 'All' && project.dataset.category !== button.dataset.filter;
    if (!project.hidden) count++;
  }
  document.querySelector('#filter-status').textContent = `${count} projects shown.`;
}
for (const button of filters) button.addEventListener('click', () => filterProjects(button));
// Keep the first view compact; filters and deep links reveal the full matching set.
if (projects.length > 9) {
  [...projects].slice(9).forEach(project => { project.hidden = true; });
  moreProjects.hidden = false;
}
moreProjects.addEventListener('click', () => {
  filterProjects(document.querySelector('[data-filter="All"]'));
  const firstRevealed = projects[9];
  if (firstRevealed) { firstRevealed.tabIndex = -1; firstRevealed.focus({preventScroll:true}); firstRevealed.scrollIntoView({block:'start'}); }
});
// Deep links from featured cards must reveal projects hidden by a previous filter.
function revealProject(hash = location.hash) {
  if (!hash.startsWith('#project-')) return;
  const target = document.getElementById(hash.slice(1));
  if (target && target.hidden) {
    filterProjects(document.querySelector('[data-filter="All"]'));
    target.scrollIntoView();
  }
}
window.addEventListener('hashchange', () => revealProject());
for (const link of document.querySelectorAll('a[href^="#project-"]')) link.addEventListener('click', () => revealProject(link.hash));
revealProject();
const tabs = [...document.querySelectorAll('[data-skill]')];
function selectSkill(tab) {
  for (const item of tabs) {
    const selected = item === tab;
    item.setAttribute('aria-selected', String(selected));
    item.tabIndex = selected ? 0 : -1;
    document.getElementById(item.getAttribute('aria-controls')).hidden = !selected;
  }
}
for (const tab of tabs) {
  tab.addEventListener('click', () => selectSkill(tab));
  tab.addEventListener('keydown', event => {
    let index = tabs.indexOf(tab);
    if (event.key === 'ArrowRight') index = (index + 1) % tabs.length;
    else if (event.key === 'ArrowLeft') index = (index + tabs.length - 1) % tabs.length;
    else if (event.key === 'Home') index = 0;
    else if (event.key === 'End') index = tabs.length - 1;
    else return;
    event.preventDefault();
    selectSkill(tabs[index]);
    tabs[index].focus();
  });
}
const menu = document.querySelector('.menu-toggle');
const nav = document.querySelector('#main-nav');
function closeMenu() { menu.setAttribute('aria-expanded', 'false'); nav.classList.remove('open'); }
menu.addEventListener('click', () => {
  const open = menu.getAttribute('aria-expanded') !== 'true';
  menu.setAttribute('aria-expanded', String(open));
  nav.classList.toggle('open', open);
});
for (const link of nav.querySelectorAll('a')) link.addEventListener('click', closeMenu);
document.addEventListener('keydown', event => {
  if (event.key === 'Escape' && menu.getAttribute('aria-expanded') === 'true') {closeMenu();menu.focus();}
});
document.addEventListener('click', event => {if (!event.target.closest('.site-header')) closeMenu();});
const navLinks = [...nav.querySelectorAll('a')];
const observer = new IntersectionObserver(entries => {
  for (const entry of entries) {
    if (!entry.isIntersecting) continue;
    for (const link of navLinks) {
      const active = link.hash === `#${entry.target.id}`;
      link.classList.toggle('active', active);
      if (active) link.setAttribute('aria-current', 'location');
      else link.removeAttribute('aria-current');
    }
  }
}, {rootMargin:'-15% 0px -55% 0px'});
for (const link of navLinks) {const section=document.querySelector(link.hash);if(section) observer.observe(section);}

// A few one-time reveals; content stays visible without animation support.
const reducedMotion = matchMedia('(prefers-reduced-motion: reduce)');
const activeReveals = new Set();
const revealTargets = document.querySelectorAll('.hero-copy, .about-section, .dashboard, .section-heading, #experience > h2, #contact > div:first-child');
if ('animate' in Element.prototype && 'IntersectionObserver' in window) {
  const reveals = new IntersectionObserver(entries => {
    for (const entry of entries) {
      if (!entry.isIntersecting) continue;
      reveals.unobserve(entry.target);
      if (reducedMotion.matches || entry.target.contains(document.activeElement)) continue;
      const animation = entry.target.animate([
        {opacity:0, transform:'translateY(12px)'},
        {opacity:1, transform:'translateY(0)'}
      ], {duration:380, easing:'cubic-bezier(.2,.7,.3,1)'});
      activeReveals.add(animation);
      animation.finished.catch(() => {}).finally(() => activeReveals.delete(animation));
    }
  }, {threshold:0.06});
  revealTargets.forEach(target => reveals.observe(target));
  reducedMotion.addEventListener('change', () => {
    if (reducedMotion.matches) activeReveals.forEach(animation => animation.cancel());
  });
  // Navigation and keyboard focus take precedence over decorative motion.
  document.addEventListener('focusin', () => activeReveals.forEach(animation => animation.cancel()));
  window.addEventListener('hashchange', () => activeReveals.forEach(animation => animation.cancel()));
}

// Keep section navigation on the clean page URL, including keyboard activation.
for (const link of document.querySelectorAll('a[href^="#"]')) {
  link.addEventListener('click', event => {
    if (event.defaultPrevented || event.button !== 0 || event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return;
    const target = document.getElementById(link.hash.slice(1));
    if (!target) return;
    event.preventDefault();
    revealProject(link.hash);
    closeMenu();
    activeReveals.forEach(animation => animation.cancel());
    if (!target.hasAttribute('tabindex')) target.tabIndex = -1;
    target.focus({preventScroll:true});
    target.scrollIntoView({block:'start', behavior:reducedMotion.matches ? 'instant' : 'smooth'});
    if (location.hash) history.replaceState(history.state, '', location.pathname + location.search);
  });
}
