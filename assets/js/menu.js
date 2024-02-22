// Mobile menu

const menuTrigger = document.querySelector(".menu-trigger");
const menu = document.querySelector(".menu");
const mobileQuery = getComputedStyle(document.body).getPropertyValue(
  "--phoneWidth"
);
const isMobile = () => window.matchMedia(mobileQuery).matches;
const isMobileMenu = () => {
  menuTrigger && menuTrigger.classList.toggle("hidden", !isMobile());
  menu && menu.classList.toggle("hidden", isMobile());
};
const showLinks = (state) => document.querySelector('body').classList.toggle('show-links', state);

isMobileMenu();

menuTrigger &&
  menuTrigger.addEventListener(
    "click",
    () => menu && menu.classList.toggle("hidden")
  );

window.addEventListener("resize", isMobileMenu);
window.addEventListener('scroll', () => {
  showLinks(true);
  
  clearTimeout(this.scrollEndTimer);
  this.scrollEndTimer = setTimeout(() => showLinks(false), 500);
}, false);
