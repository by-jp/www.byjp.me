/* Accent colour */

function updateAccentColor() {
  const now = new Date()
  let start = new Date(now)
  const hue = (now - start.setHours(0,0,0,0)) / 240000
  document.documentElement.style.setProperty('--accent', `lch(60% 120 ${hue})`);
  document.documentElement.style.setProperty('--accentAlt', `lch(60% 120 ${(hue + 120) % 360})`);
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
      const clapCount = addInteractions(btn, data.claps);
      localStorage.setItem(clapKey(btn.parentElement.action), clapCount);
      forEveryClapButton((btn) => {
        setClapCount(btn, clapCount);
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

const addInteractions = (btn, claps) => (claps + (btn.dataset.interactions || 0));

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
    btn.addEventListener("click", performClap)

    const action = btn.parentElement.action
    const lastClappedTo = localStorage.getItem(clapKey(action));
    if (lastClappedTo) {
      btn.parentElement.classList.add('clapped')
      setClapCount(btn, lastClappedTo)
    }
    toCheck.push(action)
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
        const clapCount = addInteractions(btn, data.claps);
        forEveryClapButton((btn) => {
          setClapCount(btn, clapCount);
        }, action)
      })
      .catch(console.error)
  })
})

/* drops */

const dropsContainer = document.getElementById('drops');
let drops = [];
const dropLength = 0.5; // s
const fps = 60;
const dropIncrement = 1/(fps*dropLength);
function createDrop(posX, posY) {
  startDropCounter();

  const el = document.createElement('div');
  let size = 0;

  el.style.setProperty('--dropSize', size);
  el.style.setProperty('top', posY);
  el.style.setProperty('left', posX);
  drops.push({ el, size });
  dropsContainer.appendChild(el);
  playDroplet();
}

let dropCounterInterval;
function startDropCounter() {
  if (dropCounterInterval) {
    return;
  }
  dropCounterInterval = setInterval(updateDropCounter, 1000/fps);
}

function updateDropCounter() {
  for (const [i, drop] of drops.entries()) {
    if (drop.size > 1) {
      dropsContainer.removeChild(drop.el);
      drops.splice(i, 1);
      continue;
    }

    drop.size += dropIncrement;
    drop.el.style.setProperty('--dropSize', `${drop.size}`);
  }

  if (drops.length == 0) {
    clearInterval(dropCounterInterval);
    dropCounterInterval = false;
  }
}

const audioContext = new (window.AudioContext || window.webkitAudioContext)();
function playDroplet() {
  const playLength = dropLength * 2 / 3;
  const startFreq = 98 + gRand(12);
  const endFreq = 220 + gRand(12);
  
  const oscillator = audioContext.createOscillator();
  oscillator.type = 'sine'
  oscillator.frequency.setValueAtTime(startFreq, audioContext.currentTime);
  oscillator.frequency.exponentialRampToValueAtTime(endFreq, audioContext.currentTime + playLength);
  
  const gainNode = audioContext.createGain();
  gainNode.gain.setValueAtTime(0.2, audioContext.currentTime);
  gainNode.gain.linearRampToValueAtTime(0, audioContext.currentTime + playLength);

  oscillator.connect(gainNode);
  gainNode.connect(audioContext.destination);

  oscillator.start();
  oscillator.stop(audioContext.currentTime + playLength);
}

let raining = false;
document.body.addEventListener('click', (e) => {
  raining = false;
  switch(e.target.tagName) {
    case 'A':
    case 'BUTTON':
      return true;
    default:
  }

  createDrop(`${e.pageX}px`, `${e.pageY}px`);
  return true;
});

function gRand(scale = 1) {
  return scale * Math.sqrt(-2 * Math.log(1 - Math.random())) * Math.cos(2 * Math.PI * Math.random())
}

function doRain() {
  if (!raining) {
    return
  }

  createDrop(`${Math.random() * document.body.scrollWidth}px`, `${Math.random() * document.body.scrollHeight}px`);
  setTimeout(doRain, 250+gRand(250));
}

function rain() {
  raining = true;
  doRain();
}