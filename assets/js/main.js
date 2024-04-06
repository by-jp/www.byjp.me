function updateAccentColor() {
  const now = new Date()
  let start = new Date(now)
  const hue = (now - start.setHours(0,0,0,0)) / 240000
  document.documentElement.style.setProperty('--accent', `lch(60% 120 ${hue})`);
  document.documentElement.style.setProperty('--accentHue', hue);
}
setInterval(updateAccentColor, 15000)
updateAccentColor()

const performClap = (e) => {
  e.preventDefault();
  const btn = e.currentTarget;
  btn.disabled = true;
  fetch(btn.parentElement.action, { method: 'POST', headers: new Headers({ 'Accept': 'application/json' }) })
    .then(res => res.json())
    .then(data => {
      localStorage.setItem(clapKey(btn.parentElement.action), data.claps);
      setClapCount(btn, data.claps);
      btn.parentElement.classList.add('clapped')
      btn.disabled = false;
    })
    .catch(e => console.error(`Failed to automate the clap: ${e}`))
}

const clapKey = (action) => `clap:${(new URL(action)).pathname}`;

const setClapCount = (btn, clapCount) => {
  const count = btn.querySelector('span');
  if (count) {
    count.textContent = Math.max(count.textContent, clapCount)
  } else {
    const newCount = document.createElement('span');
    newCount.textContent = clapCount;
    btn.appendChild(newCount);
  }
}

for (const btn of document.querySelectorAll('form.claps button')) {
  const lastClappedTo = localStorage.getItem(clapKey(btn.parentElement.action));
  if (lastClappedTo) {
    btn.parentElement.classList.add('clapped')
    setClapCount(btn, lastClappedTo)
  }
  btn.addEventListener("click", performClap)
}
