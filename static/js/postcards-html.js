function v() {
}
function K(e, t) {
  for (const n in t)
  e[n] = t[n];
  return e;
}
function Qt(e) {
  return e && typeof e == "object" && typeof e.then == "function";
}
function ft(e) {
  return e();
}
function bt() {
  return (/* @__PURE__ */Object.create(null));
}
function U(e) {
  e.forEach(ft);
}
function Lt(e) {
  return typeof e == "function";
}
function ut(e, t) {
  return e != e ? t == t : e !== t || e && typeof e == "object" || typeof e == "function";
}
let Y;
function te(e, t) {
  return Y || (Y = document.createElement("a")), Y.href = t, e === Y.href;
}
function Dt(e) {
  return Object.keys(e).length === 0;
}
function tt(e) {
  const t = {};
  for (const n in e)
  n[0] !== "$" && (t[n] = e[n]);
  return t;
}
let nt = !1;
function ee() {
  nt = !0;
}
function ne() {
  nt = !1;
}
function re(e, t, n, r) {
  for (; e < t;) {
    const i = e + (t - e >> 1);
    n(i) <= r ? e = i + 1 : t = i;
  }
  return e;
}
function ie(e) {
  if (e.hydrate_init)
  return;
  e.hydrate_init = !0;
  let t = e.childNodes;
  if (e.nodeName === "HEAD") {
    const l = [];
    for (let a = 0; a < t.length; a++) {
      const f = t[a];
      f.claim_order !== void 0 && l.push(f);
    }
    t = l;
  }
  const n = new Int32Array(t.length + 1),r = new Int32Array(t.length);
  n[0] = -1;
  let i = 0;
  for (let l = 0; l < t.length; l++) {
    const a = t[l].claim_order,f = (i > 0 && t[n[i]].claim_order <= a ? i + 1 : re(1, i, d => t[n[d]].claim_order, a)) - 1;
    r[l] = n[f] + 1;
    const u = f + 1;
    n[u] = l, i = Math.max(u, i);
  }
  const s = [],o = [];
  let c = t.length - 1;
  for (let l = n[i] + 1; l != 0; l = r[l - 1]) {
    for (s.push(t[l - 1]); c >= l; c--)
    o.push(t[c]);
    c--;
  }
  for (; c >= 0; c--)
  o.push(t[c]);
  s.reverse(), o.sort((l, a) => l.claim_order - a.claim_order);
  for (let l = 0, a = 0; l < o.length; l++) {
    for (; a < s.length && o[l].claim_order >= s[a].claim_order;)
    a++;
    const f = a < s.length ? s[a] : null;
    e.insertBefore(o[l], f);
  }
}
function k(e, t) {
  e.appendChild(t);
}
function B(e, t) {
  if (nt) {
    for (ie(e), (e.actual_end_child === void 0 || e.actual_end_child !== null && e.actual_end_child.parentNode !== e) && (e.actual_end_child = e.firstChild); e.actual_end_child !== null && e.actual_end_child.claim_order === void 0;)
    e.actual_end_child = e.actual_end_child.nextSibling;
    t !== e.actual_end_child ? (t.claim_order !== void 0 || t.parentNode !== e) && e.insertBefore(t, e.actual_end_child) : e.actual_end_child = t.nextSibling;
  } else
  (t.parentNode !== e || t.nextSibling !== null) && e.appendChild(t);
}
function b(e, t, n) {
  e.insertBefore(t, n || null);
}
function oe(e, t, n) {
  nt && !n ? B(e, t) : (t.parentNode !== e || t.nextSibling != n) && e.insertBefore(t, n || null);
}
function _(e) {
  e.parentNode.removeChild(e);
}
function jt(e, t) {
  for (let n = 0; n < e.length; n += 1)
  e[n] && e[n].d(t);
}
function E(e) {
  return document.createElement(e);
}
function O(e) {
  return document.createElementNS("http://www.w3.org/2000/svg", e);
}
function x(e) {
  return document.createTextNode(e);
}
function P() {
  return x(" ");
}
function rt() {
  return x("");
}
function se(e, t, n, r) {
  return e.addEventListener(t, n, r), () => e.removeEventListener(t, n, r);
}
function h(e, t, n) {
  n == null ? e.removeAttribute(t) : e.getAttribute(t) !== n && e.setAttribute(t, n);
}
function yt(e, t) {
  for (const n in t)
  h(e, n, t[n]);
}
function A(e) {
  return Array.from(e.childNodes);
}
function le(e) {
  e.claim_info === void 0 && (e.claim_info = { last_index: 0, total_claimed: 0 });
}
function ce(e, t, n, r, i = !1) {
  le(e);
  const s = (() => {
    for (let o = e.claim_info.last_index; o < e.length; o++) {
      const c = e[o];
      if (t(c)) {
        const l = n(c);
        return l === void 0 ? e.splice(o, 1) : e[o] = l, i || (e.claim_info.last_index = o), c;
      }
    }
    for (let o = e.claim_info.last_index - 1; o >= 0; o--) {
      const c = e[o];
      if (t(c)) {
        const l = n(c);
        return l === void 0 ? e.splice(o, 1) : e[o] = l, i ? l === void 0 && e.claim_info.last_index-- : e.claim_info.last_index = o, c;
      }
    }
    return r();
  })();
  return s.claim_order = e.claim_info.total_claimed, e.claim_info.total_claimed += 1, s;
}
function ae(e, t, n, r) {
  return ce(e, i => i.nodeName === t, i => {
    const s = [];
    for (let o = 0; o < i.attributes.length; o++) {
      const c = i.attributes[o];
      n[c.name] || s.push(c.name);
    }
    s.forEach(o => i.removeAttribute(o));
  }, () => r(t));
}
function M(e, t, n) {
  return ae(e, t, n, O);
}
function j(e, t) {
  t = "" + t, e.wholeText !== t && (e.data = t);
}
function m(e, t, n, r) {
  n === null ? e.style.removeProperty(t) : e.style.setProperty(t, n, r ? "important" : "");
}
function kt(e, t, n) {
  e.classList[n ? "add" : "remove"](t);
}
function fe(e, t, { bubbles: n = !1, cancelable: r = !1 } = {}) {
  const i = document.createEvent("CustomEvent");
  return i.initCustomEvent(e, n, r, t), i;
}
function Ft(e) {
  const t = {};
  for (const n of e)
  t[n.name] = n.value;
  return t;
}
let V;
function C(e) {
  V = e;
}
function dt() {
  if (!V)
  throw new Error("Function called outside component initialization");
  return V;
}
function ue() {
  const e = dt();
  return (t, n, { cancelable: r = !1 } = {}) => {
    const i = e.$$.callbacks[t];
    if (i) {
      const s = fe(t, n, { cancelable: r });
      return i.slice().forEach(o => {
        o.call(e, s);
      }), !s.defaultPrevented;
    }
    return !0;
  };
}
const I = [],wt = [],Z = [],vt = [],de = Promise.resolve();
let st = !1;
function he() {
  st || (st = !0, de.then(S));
}
function lt(e) {
  Z.push(e);
}
const ot = /* @__PURE__ */new Set();
let G = 0;
function S() {
  const e = V;
  do {
    for (; G < I.length;) {
      const t = I[G];
      G++, C(t), pe(t.$$);
    }
    for (C(null), I.length = 0, G = 0; wt.length;)
    wt.pop()();
    for (let t = 0; t < Z.length; t += 1) {
      const n = Z[t];
      ot.has(n) || (ot.add(n), n());
    }
    Z.length = 0;
  } while (I.length);
  for (; vt.length;)
  vt.pop()();
  st = !1, ot.clear(), C(e);
}
function pe(e) {
  if (e.fragment !== null) {
    e.update(), U(e.before_update);
    const t = e.dirty;
    e.dirty = [-1], e.fragment && e.fragment.p(e.ctx, t), e.after_update.forEach(lt);
  }
}
const Q = /* @__PURE__ */new Set();
let H;
function zt() {
  H = {
    r: 0,
    c: [],
    p: H };

}
function Ut() {
  H.r || U(H.c), H = H.p;
}
function T(e, t) {
  e && e.i && (Q.delete(e), e.i(t));
}
function L(e, t, n, r) {
  if (e && e.o) {
    if (Q.has(e))
    return;
    Q.add(e), H.c.push(() => {
      Q.delete(e), r && (n && e.d(1), r());
    }), e.o(t);
  } else
  r && r();
}
function Wt(e, t) {
  const n = t.token = {};
  function r(i, s, o, c) {
    if (t.token !== n)
    return;
    t.resolved = c;
    let l = t.ctx;
    o !== void 0 && (l = l.slice(), l[o] = c);
    const a = i && (t.current = i)(l);
    let f = !1;
    t.block && (t.blocks ? t.blocks.forEach((u, d) => {
      d !== s && u && (zt(), L(u, 1, 1, () => {
        t.blocks[d] === u && (t.blocks[d] = null);
      }), Ut());
    }) : t.block.d(1), a.c(), T(a, 1), a.m(t.mount(), t.anchor), f = !0), t.block = a, t.blocks && (t.blocks[s] = a), f && S();
  }
  if (Qt(e)) {
    const i = dt();
    if (e.then(s => {
      C(i), r(t.then, 1, t.value, s), C(null);
    }, s => {
      if (C(i), r(t.catch, 2, t.error, s), C(null), !t.hasCatch)
      throw s;
    }), t.current !== t.pending)
    return r(t.pending, 0), !0;
  } else {
    if (t.current !== t.then)
    return r(t.then, 1, t.value, e), !0;
    t.resolved = e;
  }
}
function qt(e, t, n) {
  const r = t.slice(),{ resolved: i } = e;
  e.current === e.then && (r[e.value] = i), e.current === e.catch && (r[e.error] = i), e.block.p(r, n);
}
function me(e, t) {
  const n = {},r = {},i = { $$scope: 1 };
  let s = e.length;
  for (; s--;) {
    const o = e[s],c = t[s];
    if (c) {
      for (const l in o)
      l in c || (r[l] = 1);
      for (const l in c)
      i[l] || (n[l] = c[l], i[l] = 1);
      e[s] = c;
    } else
    for (const l in o)
    i[l] = 1;
  }
  for (const o in r)
  o in n || (n[o] = void 0);
  return n;
}
function It(e) {
  e && e.c();
}
function ht(e, t, n, r) {
  const { fragment: i, on_mount: s, on_destroy: o, after_update: c } = e.$$;
  i && i.m(t, n), r || lt(() => {
    const l = s.map(ft).filter(Lt);
    o ? o.push(...l) : U(l), e.$$.on_mount = [];
  }), c.forEach(lt);
}
function it(e, t) {
  const n = e.$$;
  n.fragment !== null && (U(n.on_destroy), n.fragment && n.fragment.d(t), n.on_destroy = n.fragment = null, n.ctx = []);
}
function _e(e, t) {
  e.$$.dirty[0] === -1 && (I.push(e), he(), e.$$.dirty.fill(0)), e.$$.dirty[t / 31 | 0] |= 1 << t % 31;
}
function pt(e, t, n, r, i, s, o, c = [-1]) {
  const l = V;
  C(e);
  const a = e.$$ = {
    fragment: null,
    ctx: null,
    props: s,
    update: v,
    not_equal: i,
    bound: bt(),
    on_mount: [],
    on_destroy: [],
    on_disconnect: [],
    before_update: [],
    after_update: [],
    context: new Map(t.context || (l ? l.$$.context : [])),
    callbacks: bt(),
    dirty: c,
    skip_bound: !1,
    root: t.target || l.$$.root };

  o && o(a.root);
  let f = !1;
  if (a.ctx = n ? n(e, t.props || {}, (u, d, ...y) => {
    const w = y.length ? y[0] : d;
    return a.ctx && i(a.ctx[u], a.ctx[u] = w) && (!a.skip_bound && a.bound[u] && a.bound[u](w), f && _e(e, u)), d;
  }) : [], a.update(), f = !0, U(a.before_update), a.fragment = r ? r(a.ctx) : !1, t.target) {
    if (t.hydrate) {
      ee();
      const u = A(t.target);
      a.fragment && a.fragment.l(u), u.forEach(_);
    } else
    a.fragment && a.fragment.c();
    t.intro && T(e.$$.fragment), ht(e, t.target, t.anchor, t.customElement), ne(), S();
  }
  C(l);
}
let mt;
typeof HTMLElement == "function" && (mt = class extends HTMLElement {
  constructor() {
    super(), this.attachShadow({ mode: "open" });
  }
  connectedCallback() {
    const { on_mount: e } = this.$$;
    this.$$.on_disconnect = e.map(ft).filter(Lt);
    for (const t in this.$$.slotted)
    this.appendChild(this.$$.slotted[t]);
  }
  attributeChangedCallback(e, t, n) {
    this[e] = n;
  }
  disconnectedCallback() {
    U(this.$$.on_disconnect);
  }
  $destroy() {
    it(this, 1), this.$destroy = v;
  }
  $on(e, t) {
    const n = this.$$.callbacks[e] || (this.$$.callbacks[e] = []);
    return n.push(t), () => {
      const r = n.indexOf(t);
      r !== -1 && n.splice(r, 1);
    };
  }
  $set(e) {
    this.$$set && !Dt(e) && (this.$$.skip_bound = !0, this.$$set(e), this.$$.skip_bound = !1);
  }});

class ge {
  $destroy() {
    it(this, 1), this.$destroy = v;
  }
  $on(t, n) {
    const r = this.$$.callbacks[t] || (this.$$.callbacks[t] = []);
    return r.push(n), () => {
      const i = r.indexOf(n);
      i !== -1 && r.splice(i, 1);
    };
  }
  $set(t) {
    this.$$set && !Dt(t) && (this.$$.skip_bound = !0, this.$$set(t), this.$$.skip_bound = !1);
  }}

var et;
(function (e) {
  e.Book = "book", e.LeftHand = "left-hand", e.Calendar = "calendar", e.RightHand = "right-hand";
})(et || (et = {}));
const be = e => e == et.LeftHand || e == et.RightHand;
class q {
  constructor(t) {
    this.originalLocale = t.original, this.originalText = this.originalLocale && t[this.originalLocale], this._innerMap = new Map(Object.entries(t));
  }
  pickBest(t = [], n = !0) {
    if ((t.length === 0 || n) && this.originalLocale && this.originalText)
    return [this.originalText, this.originalLocale];
    if (t.length === 0) {
      const s = this._innerMap.entries(),{ value: [o, c] } = s.next();
      return [c, o];
    }
    let i = [];
    for (let s of t) {
      if (this._innerMap.has(s))
      return [this._innerMap.get(s), s];
      i.push(s.split("-")[0]);
    }
    for (let [s, o] of this._innerMap.entries())
    if (i.includes(s.split("-")[0]))
    return [o, s];
    return null;
  }}

class ye {
  constructor(t) {
    this.reader = t.getReader(), this.buffer = new Uint8Array(0);
  }
  async read(t) {
    for (; this.buffer.byteLength < t;)
    await this.addToBuffer();
    const n = this.buffer.slice(0, t);
    return this.buffer = this.buffer.slice(t), n;
  }
  async readSized() {
    const t = await this.read(4),n = new DataView(t.buffer).getUint32(0, !1);
    return this.read(n);
  }
  async addToBuffer() {
    const n = (await this.reader.read()).value;
    if (!n)
    throw "Unexpected end of data";
    const r = new Uint8Array(this.buffer.byteLength + n.byteLength);
    r.set(this.buffer, 0), r.set(n, this.buffer.byteLength), this.buffer = r;
  }}

class ke {
  constructor(t, n) {
    if (typeof t != "object" || typeof t.pxW != "number" || typeof t.pxH != "number")
    throw "Missing size attribute";
    this.isHeteroriented = be(n), this.pxW = t.pxW, this.pxH = t.pxH, this.frontAspectRatio = [this.pxW, this.pxH], t.cmW && t.cmH && (this.cmW = $t(t.cmW), this.cmH = $t(t.cmH), this.physical = `${Et(this.cmW)}cm x ${Et(this.cmH)}cm`, this.frontAspectRatio = [this.cmW[0] * this.cmH[1], this.cmW[1] * this.cmH[0]]);
    const r = Kt(this.frontAspectRatio[0], this.frontAspectRatio[1]);
    this.frontAspectRatio = [this.frontAspectRatio[0] / r, this.frontAspectRatio[1] / r];
  }
  aspectRatio(t) {
    return !this.isHeteroriented || t ? this.frontAspectRatio : [this.frontAspectRatio[1], this.frontAspectRatio[0]];
  }
  css(t) {
    if (!this.isHeteroriented)
    return {
      width: "100%",
      height: "100%",
      margin: "0 0" };

    const [n, r] = this.aspectRatio(t);
    if (n > r) {
      const i = r / n;
      return {
        width: "100%",
        height: `${100 * i}%`,
        margin: `${50 * (1 - i)}% 0` };

    } else {
      const i = n / r;
      return {
        width: `${100 * i}%`,
        height: "100%",
        margin: `0 ${50 * (1 - i)}%` };

    }
  }}

const Kt = (e, t) => t ? Kt(t, e % t) : e,we = /^(\d+)\/(\d+)$/,$t = e => {
  const t = e.match(we);
  if (!t)
  throw `invalid rational length: ${e}`;
  return [parseInt(t[1]), parseInt(t[2])];
},Et = e => (e[0] / e[1]).toFixed(1),ve = e => {
  if (e.status !== 200)
  throw "Unable to load file";
  return e;
},$e = async e => {
  const t = await fetch(e).then(d => ve(d).body);
  if (t === null)
  throw "Unable to retrieve postcard";
  const n = new ye(t);
  let r = {};
  const i = await n.read(8);
  if (new TextDecoder().decode(i) !== "postcard")
  throw "Not a postcard file";
  const s = await n.read(3),o = new DataView(s.buffer),c = o.getUint8(0),l = o.getUint8(1),a = o.getUint8(2);
  r.version = `${c}.${l}.${a}`;
  const f = await n.readSized(),u = new TextDecoder("utf-8");
  return r.metadata = Ee(JSON.parse(u.decode(f))), r.frontData = n.readSized(), r.backData = r.frontData.then(() => n.readSized()), r;
},Ee = e => {
  const { flip: t, frontSize: n, sentOn: r, front: i, back: s, context: o, ...c } = e;
  return {
    ...c,
    flip: t,
    sentOn: r && new Date(r),
    size: new ke(n, t),
    front: i && {
      description: i.description && new q(i.description),
      transcription: i.transcription && new q(i.transcription),
      secrets: i.secrets },

    back: s && {
      description: s.description && new q(s.description),
      transcription: s.transcription && new q(s.transcription),
      secrets: s.secrets },

    context: o && {
      author: o.author,
      description: o.description && new q(o.description) } };


};
function Oe(e) {
  return e && e.__esModule && Object.prototype.hasOwnProperty.call(e, "default") ? e.default : e;
}
var ct = { exports: {} },Vt = {},D = {},_t = {};
(function (e) {
  Object.defineProperty(e, "__esModule", {
    value: !0 }),
  e.default = void 0;
  function t(r, i) {
    const s = RegExp(r, "g");
    return o => {
      if (typeof o != "string")
      throw new TypeError("expected an argument of type string, but got ".concat(typeof styleObj));
      return o.match(s) ? o.replace(s, i) : o;
    };
  }
  var n = t;
  e.default = n;
})(_t);
Object.defineProperty(D, "__esModule", {
  value: !0 });

D.snakeToKebab = D.camelToKebab = void 0;
var Xt = Pe(_t);
function Pe(e) {
  return e && e.__esModule ? e : { default: e };
}
const xe = (0, Xt.default)(/[A-Z]/, e => "-".concat(e.toLowerCase()));
D.camelToKebab = xe;
const Re = (0, Xt.default)(/_/, () => "-");
D.snakeToKebab = Re;
(function (e) {
  Object.defineProperty(e, "__esModule", {
    value: !0 }),
  e.default = void 0;
  var t = D;
  function n(i, s = t.camelToKebab) {
    if (!i || typeof i != "object" || Array.isArray(i))
    throw new TypeError("expected an argument of type object, but got ".concat(typeof i));
    return Object.keys(i).map(c => "".concat(s(c), ": ").concat(i[c], ";")).join(`
`);
  }
  var r = n;
  e.default = r;
})(Vt);
(function (e, t) {
  Object.defineProperty(t, "__esModule", {
    value: !0 }),
  Object.defineProperty(t, "createParser", {
    enumerable: !0,
    get: function () {
      return r.default;
    } }),
  t.parsers = t.default = void 0;
  var n = o(Vt),r = o(_t),i = s(D);
  t.parsers = i;
  function s(f) {
    if (f && f.__esModule)
    return f;
    var u = {};
    if (f != null) {
      for (var d in f)
      if (Object.prototype.hasOwnProperty.call(f, d)) {
        var y = Object.defineProperty && Object.getOwnPropertyDescriptor ? Object.getOwnPropertyDescriptor(f, d) : {};
        y.get || y.set ? Object.defineProperty(u, d, y) : u[d] = f[d];
      }
    }
    return u.default = f, u;
  }
  function o(f) {
    return f && f.__esModule ? f : { default: f };
  }
  function c(f) {
    for (var u = 1; u < arguments.length; u++) {
      var d = arguments[u] != null ? arguments[u] : {},y = Object.keys(d);
      typeof Object.getOwnPropertySymbols == "function" && (y = y.concat(Object.getOwnPropertySymbols(d).filter(function (w) {
        return Object.getOwnPropertyDescriptor(d, w).enumerable;
      }))), y.forEach(function (w) {
        l(f, w, d[w]);
      });
    }
    return f;
  }
  function l(f, u, d) {
    return u in f ? Object.defineProperty(f, u, { value: d, enumerable: !0, configurable: !0, writable: !0 }) : f[u] = d, f;
  }
  var a = n.default;
  t.default = a, e.exports = n.default, e.exports.createParser = r.default, e.exports.parsers = c({}, i);
})(ct, ct.exports);
const Ae = /* @__PURE__ */Oe(ct.exports);
var at = /* @__PURE__ */(e => (e.FrontOnly = "front", e.BackOnly = "back", e.BothFrontFirst = "front back", e.BothBackFirst = "back front", e))(at || {});
const Ot = (e, t) => {
  const [n, r] = e.aspectRatio(t);
  return r > n;
},J = e => e.toISOString().replace(/T.*$/, ""),Ce = e => {
  let t = [];
  for (let n of e)
  t.push(`${n[0]},${n[1]}`);
  return t.join(" ");
};
function Se(e) {
  let t,n,r,i,s,o,c,l,a = [
  { width: "128" },
  { viewBox: "0 0 150.3 107.3" },
  { xmlns: "http://www.w3.org/2000/svg" },
  e[0]],
  f = {};
  for (let u = 0; u < a.length; u += 1)
  f = K(f, a[u]);
  return {
    c() {
      t = O("svg"), n = O("g"), r = O("rect"), i = O("rect"), s = O("path"), o = O("path"), c = O("path"), l = O("path"), this.h();
    },
    l(u) {
      t = M(u, "svg", { width: !0, viewBox: !0, xmlns: !0 });
      var d = A(t);
      n = M(d, "g", { transform: !0, style: !0 });
      var y = A(n);
      r = M(y, "rect", {
        style: !0,
        width: !0,
        height: !0,
        x: !0,
        y: !0,
        rx: !0,
        ry: !0 }),
      A(r).forEach(_), i = M(y, "rect", {
        style: !0,
        width: !0,
        height: !0,
        x: !0,
        y: !0,
        rx: !0,
        ry: !0 }),
      A(i).forEach(_), y.forEach(_), s = M(d, "path", { style: !0, d: !0, transform: !0 }), A(s).forEach(_), o = M(d, "path", { style: !0, d: !0, transform: !0 }), A(o).forEach(_), c = M(d, "path", { style: !0, d: !0, transform: !0 }), A(c).forEach(_), l = M(d, "path", { style: !0, d: !0, transform: !0 }), A(l).forEach(_), d.forEach(_), this.h();
    },
    h() {
      m(r, "fill", "#fff"), m(r, "stroke-width", "2.3"), m(r, "stroke-linecap", "round"), m(r, "stroke-linejoin", "round"), m(r, "stroke-dasharray", "none"), m(r, "stroke-opacity", "1"), h(r, "width", "148"), h(r, "height", "105"), h(r, "x", "51.15"), h(r, "y", "51.15"), h(r, "rx", "1"), h(r, "ry", "1"), m(i, "fill", "#fff"), m(i, "stroke-width", "2.5"), m(i, "stroke-linecap", "round"), m(i, "stroke-linejoin", "round"), m(i, "stroke-dasharray", "none"), m(i, "stroke-opacity", "1"), h(i, "width", "21"), h(i, "height", "24"), h(i, "x", "168.15"), h(i, "y", "61.15"), h(i, "rx", "1"), h(i, "ry", "1"), h(n, "transform", "translate(-50 -50)"), m(n, "fill", "#fff"), m(s, "fill", "none"), m(s, "stroke-width", "2.5"), m(s, "stroke-linecap", "round"), m(s, "stroke-linejoin", "round"), m(s, "stroke-dasharray", "none"), m(s, "stroke-opacity", "1"), h(s, "d", "m86.883 93.636 7.5 7.5m0-7.5-7.5 7.5"), h(s, "transform", "translate(3.267 -32.486)"), m(o, "fill", "none"), m(o, "stroke-width", "2.5"), m(o, "stroke-linecap", "round"), m(o, "stroke-linejoin", "round"), m(o, "stroke-dasharray", "none"), m(o, "stroke-opacity", "1"), h(o, "d", "m86.883 93.636 7.5 7.5m0-7.5-7.5 7.5"), h(o, "transform", "translate(-34.233 -32.486)"), m(c, "fill", "none"), m(c, "stroke-width", "2.5"), m(c, "stroke-linecap", "round"), m(c, "stroke-linejoin", "round"), m(c, "stroke-dasharray", "none"), m(c, "stroke-opacity", "1"), h(c, "d", "M90.633 109.317h30v3.75s0 3.75-3.75 3.75-3.75-3.75-3.75-3.75v-3.75"), h(c, "transform", "translate(-30.483 -33.167)"), m(l, "fill", "none"), m(l, "stroke-width", "2.5"), m(l, "stroke-linecap", "round"), m(l, "stroke-linejoin", "round"), m(l, "stroke-dasharray", "none"), m(l, "stroke-opacity", "1"), h(l, "d", "M125.5 67.679s3.75-3.75 7.5-3.75 7.5 3.75 7.5 3.75 3.75 3.75 7.5 3.75 7.5-3.75 7.5-3.75m-30 7.5s3.75-3.75 7.5-3.75 7.5 3.75 7.5 3.75 3.75 3.75 7.5 3.75 7.5-3.75 7.5-3.75"), h(l, "transform", "translate(-29.85 -48.279)"), yt(t, f);
    },
    m(u, d) {
      oe(u, t, d), B(t, n), B(n, r), B(n, i), B(t, s), B(t, o), B(t, c), B(t, l);
    },
    p(u, [d]) {
      yt(t, f = me(a, [
      { width: "128" },
      { viewBox: "0 0 150.3 107.3" },
      { xmlns: "http://www.w3.org/2000/svg" },
      d & 1 && u[0]]));

    },
    i: v,
    o: v,
    d(u) {
      u && _(t);
    } };

}
function Te(e, t, n) {
  return e.$$set = r => {
    n(0, t = K(K({}, t), tt(r)));
  }, t = tt(t), [t];
}
class Yt extends ge {
  constructor(t) {
    super(), pt(this, t, Te, Se, ut, {});
  }}

function Pt(e, t, n) {
  const r = e.slice();
  return r[21] = t[n], r[23] = n, r;
}
function xt(e, t, n) {
  const r = e.slice();
  return r[25] = t[n], r;
}
function Ne(e) {
  let t, n, r, i;
  return r = new Yt({}), {
    c() {
      t = E("div"), n = E("div"), It(r.$$.fragment), h(n, "class", "error"), h(t, "class", "postcard error");
    },
    m(s, o) {
      b(s, t, o), k(t, n), ht(r, n, null), i = !0;
    },
    p: v,
    i(s) {
      i || (T(r.$$.fragment, s), i = !0);
    },
    o(s) {
      L(r.$$.fragment, s), i = !1;
    },
    d(s) {
      s && _(t), it(r);
    } };

}
function Me(e) {
  let t,n,r = e[2],i = [];
  for (let o = 0; o < r.length; o += 1)
  i[o] = At(Pt(e, r, o));
  const s = o => L(i[o], 1, 1, () => {
    i[o] = null;
  });
  return {
    c() {
      t = E("div");
      for (let o = 0; o < i.length; o += 1)
      i[o].c();
      h(t, "class", "postcard flip-" + e[20].flip), kt(t, "flipped", e[0]);
    },
    m(o, c) {
      b(o, t, c);
      for (let l = 0; l < i.length; l += 1)
      i[l].m(t, null);
      n = !0;
    },
    p(o, c) {
      if (c & 62) {
        r = o[2];
        let l;
        for (l = 0; l < r.length; l += 1) {
          const a = Pt(o, r, l);
          i[l] ? (i[l].p(a, c), T(i[l], 1)) : (i[l] = At(a), i[l].c(), T(i[l], 1), i[l].m(t, null));
        }
        for (zt(), l = r.length; l < i.length; l += 1)
        s(l);
        Ut();
      }
      (!n || c & 1) && kt(t, "flipped", o[0]);
    },
    i(o) {
      if (!n) {
        for (let c = 0; c < r.length; c += 1)
        T(i[c]);
        n = !0;
      }
    },
    o(o) {
      i = i.filter(Boolean);
      for (let c = 0; c < i.length; c += 1)
      L(i[c]);
      n = !1;
    },
    d(o) {
      o && _(t), jt(i, o);
    } };

}
function Be(e) {
  let t, n, r;
  return n = new Yt({}), {
    c() {
      t = E("div"), It(n.$$.fragment), h(t, "class", "error");
    },
    m(i, s) {
      b(i, t, s), ht(n, t, null), r = !0;
    },
    p: v,
    i(i) {
      r || (T(n.$$.fragment, i), r = !0);
    },
    o(i) {
      L(n.$$.fragment, i), r = !1;
    },
    d(i) {
      i && _(t), it(n);
    } };

}
function He(e) {
  let t,n,r,i,s = e[24].secrets.length > 0 && Le(e);
  return {
    c() {
      t = E("img"), r = P(), s && s.c(), i = rt(), te(t.src, n = e[24].src) || h(t, "src", n), h(t, "alt", e[24].description), h(t, "lang", e[20].locale);
    },
    m(o, c) {
      b(o, t, c), b(o, r, c), s && s.m(o, c), b(o, i, c);
    },
    p(o, c) {
      o[24].secrets.length > 0 && s.p(o, c);
    },
    i: v,
    o: v,
    d(o) {
      o && _(t), o && _(r), s && s.d(o), o && _(i);
    } };

}
function Le(e) {
  let t,n,r,i,s,o,c,l = e[24].secrets,a = [];
  for (let f = 0; f < l.length; f += 1)
  a[f] = Rt(xt(e, l, f));
  return {
    c() {
      t = O("svg"), n = O("defs"), r = O("linearGradient"), i = O("stop"), s = O("stop"), o = O("animate"), c = O("animate");
      for (let f = 0; f < a.length; f += 1)
      a[f].c();
      h(i, "offset", "0%"), h(i, "stop-color", "rgba(255,255,255,0.2)"), h(s, "offset", "100%"), h(s, "stop-color", "rgba(200,200,200,0.4)"), h(o, "attributeName", "x1"), h(o, "dur", "700ms"), h(o, "from", "0"), h(o, "to", "0.04"), h(o, "repeatCount", "indefinite"), h(c, "attributeName", "x2"), h(c, "dur", "700ms"), h(c, "from", "0.01"), h(c, "to", "0.05"), h(c, "repeatCount", "indefinite"), h(r, "id", "secret"), h(r, "x1", "0"), h(r, "x2", "0.01"), h(r, "y1", "0.01"), h(r, "y2", "0"), h(r, "gradientUnits", "userSpaceOnUse"), h(r, "spreadMethod", "reflect"), h(r, "vector-effect", "non-scaling-size"), h(t, "class", "secrets"), h(t, "viewBox", "0 0 1 1");
    },
    m(f, u) {
      b(f, t, u), k(t, n), k(n, r), k(r, i), k(r, s), k(r, o), k(r, c);
      for (let d = 0; d < a.length; d += 1)
      a[d].m(t, null);
    },
    p(f, u) {
      if (u & 4) {
        l = f[24].secrets;
        let d;
        for (d = 0; d < l.length; d += 1) {
          const y = xt(f, l, d);
          a[d] ? a[d].p(y, u) : (a[d] = Rt(y), a[d].c(), a[d].m(t, null));
        }
        for (; d < a.length; d += 1)
        a[d].d(1);
        a.length = l.length;
      }
    },
    d(f) {
      f && _(t), jt(a, f);
    } };

}
function Rt(e) {
  let t;
  return {
    c() {
      t = O("polygon"), h(t, "points", Ce(e[25])), m(t, "fill", "url(#secret)"), m(t, "vector-effect", "non-scaling-size");
    },
    m(n, r) {
      b(n, t, r);
    },
    p: v,
    d(n) {
      n && _(t);
    } };

}
function De(e) {
  return {
    c: v,
    m: v,
    p: v,
    i: v,
    o: v,
    d: v };

}
function At(e) {
  let t,n,r,i,s,o = {
    ctx: e,
    current: null,
    token: null,
    hasCatch: !0,
    pending: De,
    then: He,
    catch: Be,
    value: 24,
    error: 28,
    blocks: [,,,] };

  return Wt(e[21], o), {
    c() {
      t = E("div"), o.block.c(), n = P(), h(t, "class", "side" + (e[4] ? " flippable" : "")), h(t, "style", Ae(e[20].size.css(e[3](e[23]))));
    },
    m(c, l) {
      b(c, t, l), o.block.m(t, o.anchor = null), o.mount = () => t, o.anchor = n, k(t, n), r = !0, i || (s = se(t, "click", e[5]), i = !0);
    },
    p(c, l) {
      e = c, qt(o, e, l);
    },
    i(c) {
      r || (T(o.block), r = !0);
    },
    o(c) {
      for (let l = 0; l < 3; l += 1) {
        const a = o.blocks[l];
        L(a);
      }
      r = !1;
    },
    d(c) {
      c && _(t), o.block.d(), o.token = null, o = null, i = !1, s();
    } };

}
function je(e) {
  let t;
  return {
    c() {
      t = E("div"), h(t, "class", "postcard loading");
    },
    m(n, r) {
      b(n, t, r);
    },
    p: v,
    i: v,
    o: v,
    d(n) {
      n && _(t);
    } };

}
function Fe(e) {
  let t,n,r = {
    ctx: e,
    current: null,
    token: null,
    hasCatch: !0,
    pending: je,
    then: Me,
    catch: Ne,
    value: 20,
    error: 28,
    blocks: [,,,] };

  return Wt(e[1], r), {
    c() {
      t = rt(), r.block.c(), this.c = v;
    },
    m(i, s) {
      b(i, t, s), r.block.m(i, r.anchor = s), r.mount = () => t.parentNode, r.anchor = t, n = !0;
    },
    p(i, [s]) {
      e = i, qt(r, e, s);
    },
    i(i) {
      n || (T(r.block), n = !0);
    },
    o(i) {
      for (let s = 0; s < 3; s += 1) {
        const o = r.blocks[s];
        L(o);
      }
      n = !1;
    },
    d(i) {
      i && _(t), r.block.d(i), r.token = null, r = null;
    } };

}
function ze(e, t, n) {
  let { src: r } = t,{ name: i = void 0 } = t,{ show: s = at.BothFrontFirst } = t;
  const o = i && document.querySelectorAll(`postcard-display[name="${i}"`).length;
  o && o > 1 && console.warn(`There are ${o} postcards with name="${i}". Any <postcard-info for="${i}"> elements will be attached to the first one only.`);
  const c = s.split(" ");
  let l = !1;
  const a = dt(),f = ue(),u = (g, R) => {
    f(g, R), a.dispatchEvent && a.dispatchEvent(new CustomEvent(g, { bubbles: !0, composed: !0, detail: R }));
  },d = g => URL.createObjectURL(new Blob([g])),y = g => {
    throw console.error(g), g;
  },w = g => {
    const [R, z] = g.size.aspectRatio(!0),Jt = Ot(g.size, !0) && g.size.isHeteroriented,Zt = Ot(g.size, !1) && g.size.isHeteroriented;
    a.style.setProperty("--frontAR", Jt ? "1/1" : `${R}/${z}`), a.style.setProperty("--backAR", Zt ? "1/1" : `${R}/${z}`), N();
  },N = () => a.style.aspectRatio = W() ? "var(--frontAR)" : "var(--backAR)",p = $e(r),$ = p.then(({ metadata: g }) => g).catch(y);
  $.then(g => {
    w(g), u("postcard-loaded", {
      name: i,
      metadata: g,
      showingSide: F() });

  });
  const X = c.map(g => p.then(R => R[`${g}Data`].then(z => ({ metadata: R.metadata, bytes: z }))).then(({ metadata: R, bytes: z }) => ({
    src: d(z),
    secrets: R[g].secrets || [],
    description: R[g].description })).
  catch(y)),F = (g = l ? 1 : 0) => c[g],W = (g = l ? 1 : 0) => F(g) === at.FrontOnly,gt = c.length === 2,Gt = gt && (() => {
    n(0, l = !l), N(), u("postcard-flipped", { name: i, showingSide: F() });
  });
  return e.$$set = g => {
    "src" in g && n(6, r = g.src), "name" in g && n(7, i = g.name), "show" in g && n(8, s = g.show);
  }, [l, $, X, W, gt, Gt, r, i, s];
}
class Ue extends mt {
  constructor(t) {
    super(), this.shadowRoot.innerHTML = "<style>:host{display:block;width:100%;height:100%;transition:aspect-ratio 1s ease-out}.postcard{position:relative;perspective:1000px;display:inline-block;width:100%}.postcard.flip-book,.postcard.flip-calendar{height:100%}.postcard.flip-right-hand,.postcard.flip-left-hand{aspect-ratio:1/1}.side{width:100%;height:100%;position:absolute;backface-visibility:hidden;transition:transform 1s ease-out;transform-style:preserve-3d;filter:drop-shadow(5px 5px 5px rgba(0, 0, 0, 0.45))}.side.flippable{cursor:pointer}.side>*{width:100%;height:100%}.side .secrets{position:absolute;top:0;left:0}.side .secrets polygon{opacity:0;transition:opacity 0.2s ease-in-out}.side .secrets polygon:hover{opacity:1}.side .error{box-sizing:border-box;width:100%;height:100%;border-radius:2%;border:2px dashed rgb(213, 161, 161);background-color:#f8f7f6;padding:2% 4%;display:flex;align-items:center;justify-content:center;flex-direction:column}.side .error svg{max-width:128px;stroke:rgb(213, 161, 161)}.postcard.flipped.flip-book .side:first-child{transform:rotateY(180deg)}.postcard:not(.flipped).flip-book .side:nth-child(2){transform:rotateY(180deg)}.postcard.flipped.flip-book .side:nth-child(2){transform:rotateY(360deg)}.postcard.flipped.flip-left-hand .side:first-child{transform:rotate3d(1, -1, 0, 180deg)}.postcard:not(.flipped).flip-left-hand .side:nth-child(2){transform:rotate3d(1, -1, 0, -180deg)}.postcard.flipped.flip-left-hand .side:nth-child(2){transform:rotate3d(1, -1, 0, 0deg)}.postcard.flipped.flip-calendar .side:first-child{transform:rotateX(180deg)}.postcard:not(.flipped).flip-calendar .side:nth-child(2){transform:rotateX(180deg)}.postcard.flipped.flip-calendar .side:nth-child(2){transform:rotateX(360deg)}.postcard.flipped.flip-right-hand .side:first-child{transform:rotate3d(1, 1, 0, 180deg)}.postcard:not(.flipped).flip-right-hand .side:nth-child(2){transform:rotate3d(1, 1, 0, -180deg)}.postcard.flipped.flip-right-hand .side:nth-child(2){transform:rotate3d(1, 1, 0, 0deg)}</style>", pt(
    this,
    {
      target: this.shadowRoot,
      props: Ft(this.attributes),
      customElement: !0 },

    ze,
    Fe,
    ut,
    { src: 6, name: 7, show: 8 },
    null),
    t && (t.target && b(t.target, this, t.anchor), t.props && (this.$set(t.props), S()));
  }
  static get observedAttributes() {
    return ["src", "name", "show"];
  }
  get src() {
    return this.$$.ctx[6];
  }
  set src(t) {
    this.$$set({ src: t }), S();
  }
  get name() {
    return this.$$.ctx[7];
  }
  set name(t) {
    this.$$set({ name: t }), S();
  }
  get show() {
    return this.$$.ctx[8];
  }
  set show(t) {
    this.$$set({ show: t }), S();
  }}

customElements.define("postcard-display", Ue);
function Ct(e) {
  let t, n, r, i, s;
  return {
    c() {
      t = E("p"), n = E("strong"), n.textContent = "Context:", r = P(), i = x(e[5]), h(t, "lang", s = e[0].locale);
    },
    m(o, c) {
      b(o, t, c), k(t, n), k(t, r), k(t, i);
    },
    p(o, c) {
      c & 32 && j(i, o[5]), c & 1 && s !== (s = o[0].locale) && h(t, "lang", s);
    },
    d(o) {
      o && _(t);
    } };

}
function St(e) {
  let t, n, r, i;
  return {
    c() {
      t = E("p"), n = E("strong"), n.textContent = "Physical size (front):", r = P(), i = x(e[4]);
    },
    m(s, o) {
      b(s, t, o), k(t, n), k(t, r), k(t, i);
    },
    p(s, o) {
      o & 16 && j(i, s[4]);
    },
    d(s) {
      s && _(t);
    } };

}
function Tt(e) {
  var l;
  let t,n,r,i,s = J((l = e[0]) == null ? void 0 : l.sentOn) + "",o,c;
  return {
    c() {
      var a;
      t = E("p"), n = E("strong"), n.textContent = "Date:", r = P(), i = E("time"), o = x(s), h(i, "datetime", c = J((a = e[0]) == null ? void 0 : a.sentOn));
    },
    m(a, f) {
      b(a, t, f), k(t, n), k(t, r), k(t, i), k(i, o);
    },
    p(a, f) {
      var u, d;
      f & 1 && s !== (s = J((u = a[0]) == null ? void 0 : u.sentOn) + "") && j(o, s), f & 1 && c !== (c = J((d = a[0]) == null ? void 0 : d.sentOn)) && h(i, "datetime", c);
    },
    d(a) {
      a && _(t);
    } };

}
function Nt(e) {
  let t,n,r,i = e[0].location.name + "",s;
  return {
    c() {
      t = E("p"), n = E("strong"), n.textContent = "Location:", r = P(), s = x(i);
    },
    m(o, c) {
      b(o, t, c), k(t, n), k(t, r), k(t, s);
    },
    p(o, c) {
      c & 1 && i !== (i = o[0].location.name + "") && j(s, i);
    },
    d(o) {
      o && _(t);
    } };

}
function Mt(e) {
  let t,n,r,i,s,o = e[3] && Bt(e),c = e[2] && Ht(e);
  return {
    c() {
      t = x("On the "), n = x(e[1]), r = x(`:

  `), o && o.c(), i = P(), c && c.c(), s = rt();
    },
    m(l, a) {
      b(l, t, a), b(l, n, a), b(l, r, a), o && o.m(l, a), b(l, i, a), c && c.m(l, a), b(l, s, a);
    },
    p(l, a) {
      a & 2 && j(n, l[1]), l[3] ? o ? o.p(l, a) : (o = Bt(l), o.c(), o.m(i.parentNode, i)) : o && (o.d(1), o = null), l[2] ? c ? c.p(l, a) : (c = Ht(l), c.c(), c.m(s.parentNode, s)) : c && (c.d(1), c = null);
    },
    d(l) {
      l && _(t), l && _(n), l && _(r), o && o.d(l), l && _(i), c && c.d(l), l && _(s);
    } };

}
function Bt(e) {
  let t, n, r, i, s;
  return {
    c() {
      t = E("p"), n = E("strong"), n.textContent = "Description:", r = P(), i = x(e[3]), h(t, "lang", s = e[0].locale);
    },
    m(o, c) {
      b(o, t, c), k(t, n), k(t, r), k(t, i);
    },
    p(o, c) {
      c & 8 && j(i, o[3]), c & 1 && s !== (s = o[0].locale) && h(t, "lang", s);
    },
    d(o) {
      o && _(t);
    } };

}
function Ht(e) {
  let t, n, r, i, s;
  return {
    c() {
      t = E("p"), n = E("strong"), n.textContent = "Transcription:", r = P(), i = x(e[2]), h(t, "lang", s = e[0].locale);
    },
    m(o, c) {
      b(o, t, c), k(t, n), k(t, r), k(t, i);
    },
    p(o, c) {
      c & 4 && j(i, o[2]), c & 1 && s !== (s = o[0].locale) && h(t, "lang", s);
    },
    d(o) {
      o && _(t);
    } };

}
function We(e) {
  let t, n;
  return {
    c() {
      var r;
      t = E("a"), n = x("Download postcard"), h(t, "href", (r = e[7]) == null ? void 0 : r.getAttribute("src"));
    },
    m(r, i) {
      b(r, t, i), k(t, n);
    },
    p: v,
    d(r) {
      r && _(t);
    } };

}
function qe(e) {
  var y, w, N;
  let t,n,r,i,s,o,c = e[5] && Ct(e),l = e[4] && St(e),a = ((y = e[0]) == null ? void 0 : y.sentOn) && Tt(e),f = ((N = (w = e[0]) == null ? void 0 : w.location) == null ? void 0 : N.name) && Nt(e),u = (e[3] || e[2]) && Mt(e),d = e[6] && We(e);
  return {
    c() {
      c && c.c(), t = P(), l && l.c(), n = P(), a && a.c(), r = P(), f && f.c(), i = P(), u && u.c(), s = P(), d && d.c(), o = rt(), this.c = v;
    },
    m(p, $) {
      c && c.m(p, $), b(p, t, $), l && l.m(p, $), b(p, n, $), a && a.m(p, $), b(p, r, $), f && f.m(p, $), b(p, i, $), u && u.m(p, $), b(p, s, $), d && d.m(p, $), b(p, o, $);
    },
    p(p, [$]) {
      var X, F, W;
      p[5] ? c ? c.p(p, $) : (c = Ct(p), c.c(), c.m(t.parentNode, t)) : c && (c.d(1), c = null), p[4] ? l ? l.p(p, $) : (l = St(p), l.c(), l.m(n.parentNode, n)) : l && (l.d(1), l = null), (X = p[0]) != null && X.sentOn ? a ? a.p(p, $) : (a = Tt(p), a.c(), a.m(r.parentNode, r)) : a && (a.d(1), a = null), (W = (F = p[0]) == null ? void 0 : F.location) != null && W.name ? f ? f.p(p, $) : (f = Nt(p), f.c(), f.m(i.parentNode, i)) : f && (f.d(1), f = null), p[3] || p[2] ? u ? u.p(p, $) : (u = Mt(p), u.c(), u.m(s.parentNode, s)) : u && (u.d(1), u = null), p[6] && d.p(p, $);
    },
    i: v,
    o: v,
    d(p) {
      c && c.d(p), p && _(t), l && l.d(p), p && _(n), a && a.d(p), p && _(r), f && f.d(p), p && _(i), u && u.d(p), p && _(s), d && d.d(p), p && _(o);
    } };

}
function Ie(e, t, n) {
  let r,i,s,o,c,{ downloadable: l = !1 } = t,{ for: a } = t;
  const f = typeof l == "boolean" ? l : !0,u = document.querySelector(`postcard-display[name="${a}"`);
  u || console.error(`No element <postcard-display name="${a}"> found to attach to from <postcard-info for="${a}">.`);
  let d, y;
  return u == null || u.addEventListener("postcard-loaded", w => {
    n(0, d = w.detail.metadata), n(1, y = w.detail.showingSide);
  }), u == null || u.addEventListener("postcard-flipped", w => n(1, y = w.detail.showingSide)), e.$$set = w => {
    n(11, t = K(K({}, t), tt(w))), "downloadable" in w && n(8, l = w.downloadable);
  }, e.$$.update = () => {
    var w, N;
    e.$$.dirty & 1 && n(5, r = ((w = d == null ? void 0 : d.context) == null ? void 0 : w.description) && d.context.description), e.$$.dirty & 1 && n(4, i = (N = d == null ? void 0 : d.size) == null ? void 0 : N.physical), e.$$.dirty & 3 && n(9, s = d && d[y]), e.$$.dirty & 512 && n(3, o = (s == null ? void 0 : s.description) && s.description), e.$$.dirty & 512 && n(2, c = (s == null ? void 0 : s.transcription) && s.transcription);
  }, t = tt(t), [
  d,
  y,
  c,
  o,
  i,
  r,
  f,
  u,
  l,
  s];

}
class Ke extends mt {
  constructor(t) {
    super(), pt(
    this,
    {
      target: this.shadowRoot,
      props: Ft(this.attributes),
      customElement: !0 },

    Ie,
    qe,
    ut,
    { downloadable: 8 },
    null),
    t && (t.target && b(t.target, this, t.anchor), t.props && (this.$set(t.props), S()));
  }
  static get observedAttributes() {
    return ["downloadable"];
  }
  get downloadable() {
    return this.$$.ctx[8];
  }
  set downloadable(t) {
    this.$$set({ downloadable: t }), S();
  }}

customElements.define("postcard-info", Ke);