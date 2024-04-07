/* Accent colour */

function updateAccentColor() {
  const now = new Date()
  let start = new Date(now)
  const hue = (now - start.setHours(0,0,0,0)) / 240000
  document.documentElement.style.setProperty('--accent', `lch(60% 120 ${hue})`);
  document.documentElement.style.setProperty('--accentHue', hue);
}
setInterval(updateAccentColor, 15000)
updateAccentColor()

/* Clap counter */

const performClap = (e) => {
  e.preventDefault();
  const btn = e.currentTarget;
  btn.disabled = true;
  fetch(btn.parentElement.action, {
    method: 'POST',
    headers: new Headers({ 'Accept': 'application/json' }),
    signal: AbortSignal.timeout(2000),
  }).then(res => res.json().then((data) => ([res, data])))
    .then(([res, data]) => {
      if (res.status !== 200) {
        const errorDesc = data.error || JSON.stringify(data)
        throw new Error(`Received HTTP ${res.status} from clap counter: ${errorDesc}`)
      }
      if (typeof data.claps !== 'number') {
        throw new Error(`Clap count response isn't a number: ${JSON.stringify(data)}`)
      }
      return data;
    })
    .then(data => {
      localStorage.setItem(clapKey(btn.parentElement.action), data.claps);
      forEveryClapButton((btn) => {
        setClapCount(btn, data.claps);
      }, btn.parentElement.action)
      btn.parentElement.classList.add('clapped')
      btn.disabled = false;
    })
    .catch(e => {
      if (e.name === 'AbortError') {
        e = new Error("Fetch exceeded timeout")
      }

      console.error(`Failed to automate the clap: ${e.message}`)
      alert("Sorry friend! It looks like my clap-o-meter has broken on us.\n\n"
      + "Instead, you could reach out to me via the sites listed on my homepage — to share appreciation for this post, or to tell me about this break!\n\n"
      + "(And if you're technically inclined,the console will list the error's details!)")
    })
}

const clapKey = (action) => `clap:${(new URL(action)).pathname}`;

const setClapCount = (btn, clapCount) => {
  if (clapCount === 0) {
    return
  }

  const count = clapCountEl(btn);
  if (count) {
    count.textContent = Math.max(count.textContent, clapCount)
  } else {
    const newCount = document.createElement('span');
    newCount.textContent = clapCount;
    btn.appendChild(newCount);
  }
}

const clapCountEl = (btn) => btn.querySelector('span')

const forEveryClapButton = (fn, sameAction = '') => {
  for (const btn of document.querySelectorAll(`form.claps${sameAction ? `[action="${sameAction}"]` : ''} button`)) {
    fn(btn)
  }
}

document.addEventListener("DOMContentLoaded", () => {
  let toCheck = []
  forEveryClapButton((btn) => {
    const action = btn.parentElement.action
    const lastClappedTo = localStorage.getItem(clapKey(action));
    toCheck.push(action)
    if (lastClappedTo) {
      btn.parentElement.classList.add('clapped')
      setClapCount(btn, lastClappedTo)
    }
    btn.addEventListener("click", performClap)
  })

  toCheck.forEach((action) => {
    fetch(action, { method: 'GET', headers: new Headers({ 'Accept': 'application/json' }) })
      .then(res => {
        if (res.status !== 200) {
          throw new Error(`Got HTTP ${res.status} while trying to retrieve claps`)
        }
        return res.json() 
      })
      .then(data => {
        forEveryClapButton((btn) => {
          setClapCount(btn, data.claps);
        }, action)
      })
      .catch(console.error)
  })
})
