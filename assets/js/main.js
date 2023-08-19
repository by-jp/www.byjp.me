function updateAccentColor() {
  const now = new Date()
  let start = new Date(now)
  const hue = (now - start.setHours(0,0,0,0)) / 240000
  document.documentElement.style.setProperty('--accent', `lch(60% 120 ${hue})`);
  document.documentElement.style.setProperty('--accentHue', hue);
}
setInterval(updateAccentColor, 15000)
updateAccentColor()
