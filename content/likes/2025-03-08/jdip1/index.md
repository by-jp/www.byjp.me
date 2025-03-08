---
date: 2025-03-08T16:57:50.992Z
publishDate: 2025-03-08T16:57:50.992Z
likeOf: https://jpcamara.com/2024/12/01/speeding-up-ruby.html
references:
  https://jpcamaraCom/2024/12/01/speedingUpRubyHtml:
    url: https://jpcamara.com/2024/12/01/speeding-up-ruby.html
    type: entry
    published: Dec 4, 2024
    name: Speeding up Ruby by rewriting C… in Ruby
    content:
      html: |-
        <p><img src="https://cdn.uploads.micro.blog/98548/2024/yjitvsc.drawio.png" alt=""></p>
        <p>There is a recent <a href="https://github.com/bddicken/languages">language comparison repo</a> which has been getting shared a lot. In it, CRuby was the third slowest option, only beating out R and Python.</p>
        <p>The repo author, <a href="https://x.com/BenjDicken">@BenjDicken</a>, <a href="https://x.com/BenjDicken/status/1861072804239847914">created a fun visualization</a> of each language’s performance. Here’s one of the visualizations, which shows Ruby as the third slowest language benchmarked:</p>
        <div id="original_visual"></div>
        <blockquote>
        <p>The code for this visualization is from <a href="https://benjdd.com/languages/">https://benjdd.com/languages/</a>, with permission from <a href="https://x.com/BenjDicken/status/1862623583803253149">@BenjDicken</a></p>
        </blockquote>
        <p>The repository describes itself as:</p>
        <blockquote>
        <p>A repo for collaboratively building small benchmarks to compare languages.</p>
        </blockquote>
        <p>It contains two different benchmarks:</p>
        <ol>
        <li>“Loops”, which “Emphasizes loop, conditional, and basic math performance”</li>
        <li>“Fibonacci”, which “Emphasizes function call overhead and recursion.”</li>
        </ol>
        <p>The loop example iterates 1 billion times, utilizing a nested loop:</p>
        <div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-ruby" data-lang="ruby">u <span style="color:#f92672">=</span> <span style="color:#66d9ef">ARGV</span><span style="color:#f92672">[</span><span style="color:#ae81ff">0</span><span style="color:#f92672">].</span>to_i       
        r <span style="color:#f92672">=</span> rand(<span style="color:#ae81ff">10_000</span>)                          
        a <span style="color:#f92672">=</span> Array<span style="color:#f92672">.</span>new(<span style="color:#ae81ff">10_000</span>, <span style="color:#ae81ff">0</span>)                 
        	
        (<span style="color:#ae81ff">0</span><span style="color:#f92672">...</span><span style="color:#ae81ff">10_000</span>)<span style="color:#f92672">.</span>each <span style="color:#66d9ef">do</span> <span style="color:#f92672">|</span>i<span style="color:#f92672">|</span>                     
          (<span style="color:#ae81ff">0</span><span style="color:#f92672">...</span><span style="color:#ae81ff">100_000</span>)<span style="color:#f92672">.</span>each <span style="color:#66d9ef">do</span> <span style="color:#f92672">|</span>j<span style="color:#f92672">|</span>               
            a<span style="color:#f92672">[</span>i<span style="color:#f92672">]</span> <span style="color:#f92672">+=</span> j <span style="color:#f92672">%</span> u                     
          <span style="color:#66d9ef">end</span>
          a<span style="color:#f92672">[</span>i<span style="color:#f92672">]</span> <span style="color:#f92672">+=</span> r                      
        <span style="color:#66d9ef">end</span>
        	
        puts a<span style="color:#f92672">[</span>r<span style="color:#f92672">]</span>
        </code></pre></div><p>The Fibonacci example is a basic “naive” Fibonacci implementation<sup id="fnref:1"><a href="#fn:1" class="footnote-ref" role="doc-noteref">1</a></sup>:</p>
        <pre><code>def fibonacci(n)
          return 0 if n == 0
          return 1 if n == 1
          fibonacci(n - 1) + fibonacci(n - 2)
        end

        u = ARGV[0].to_i
        r = 0

        (1...u).each do |i|
          r += fibonacci(i)
        end

        puts r
        </code></pre>
        <p>Run on <a href="https://x.com/BenjDicken">@BenjDicken</a>’s M3 MacBook Pro, Ruby 3.3.6 takes 28 seconds to run the loop iteration example, and 12 seconds to run the Fibonacci example. For comparison, node.js takes a little over a second for both examples - it’s not a great showing for Ruby.</p>
        <table>
          <thead>
            <tr>
              <th></th>
              <th>Fibonacci</th>
              <th>Loops</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>Ruby</td>
              <td>12.17s</td>
              <td>28.80s</td>
            </tr>
            <tr>
              <td>node.js</td>
              <td>1.11s</td>
              <td>1.03s</td>
            </tr>
          </tbody>
        </table>
        <p>From this point on, I’ll use benchmarks relative to my own computer. Running the same benchmark on my M2 MacBook Air, I get 33.43 seconds for the loops and 16.33 seconds for fibonacci - even worse 🥺. Node runs a little over 1 second for fibonacci and 2 seconds for the loop example.</p>
        <table>
          <thead>
            <tr>
              <th></th>
              <th>Fibonacci</th>
              <th>Loops</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>Ruby</td>
              <td>16.33s</td>
              <td>33.43s</td>
            </tr>
            <tr>
              <td>node.js</td>
              <td>1.36s</td>
              <td>2.07s</td>
            </tr>
          </tbody>
        </table>
        <h3 id="who-cares">Who cares?</h3>
        <p>In most ways, these types of benchmarks are meaningless. Python was the slowest language in the benchmark, and yet at the same time it’s the <a href="https://github.blog/news-insights/octoverse/octoverse-2024/">most used language on Github as of October 2024</a>. Ruby runs some of the <a href="https://x.com/tobi/status/1863935229620363693">largest web apps in the world</a>. I ran a <a href="https://x.com/jpcamara/status/1849984009515966958">benchmark recently of websocket performance between the Ruby Falcon web server and node.js</a>, and the Ruby results were close to the node.js results. Are you doing a billion loop iterations or using web sockets?</p>
        <p>A programming language should be reasonably efficient - after that the usefulness of the language, the type of tasks you work on, and language productivity outweigh the speed at which you can run a billion iterations of a loop, or complete an intentionally inefficient implementation of a Fibonacci method.</p>
        <p>That said:</p>
        <ol>
        <li>The programming world loves microbenchmarks 🤷‍♂️</li>
        <li>Having a fast benchmark may not be valuable in practice but it has meaning for people’s interest in a language. Some would claim it means you’ll have an easier time scaling performance, but that’s arguable<sup id="fnref:2"><a href="#fn:2" class="footnote-ref" role="doc-noteref">2</a></sup></li>
        <li>It’s disappointing if your language of choice doesn’t perform well. It’s nice to be able to say “I use and enjoy this language, and it runs fast in all benchmarks!”</li>
        </ol>
        <p>In the case of this Ruby benchmark, I had a feeling that YJIT wasn’t being applied in the Ruby code, so I checked the repo. Lo and behold, the command was as follows:</p>
        <pre><code>ruby ./code.rb 40
        </code></pre>
        <p>We know my results from earlier (33 seconds and 16 seconds). What do we get with YJIT applied?</p>
        <pre><code>ruby --yjit ./code.rb 40
        </code></pre>
        <table>
          <thead>
            <tr>
              <th></th>
        <th>Fibonacci</th>
        <th>Loops</th>
        </tr>
        </thead>
        <tbody>
        <tr>
        <td>Ruby</td>
        <td>2.06s</td>
        <td>25.57s</td>
        </tr>
        </tbody>
        </table>
        <p>Nice! With YJIT, Fibonacci gets a massive boost - going from 16.88 seconds down to 2.06 seconds. It’s close to the speed of node.js at that point!</p>
        <p>YJIT makes a more modest difference for the looping example - going from 33.43 seconds down to 25.57 seconds. Why is that?</p>
        <h3 id="a-team-effort">A team effort</h3>
        <p>I wasn’t alone in trying out these code samples with YJIT. On twitter, <a href="https://x.com/bsilva96">@bsilva96</a> had asked the same questions:</p>
        <p><img src="https://cdn.uploads.micro.blog/98548/2024/screenshot-2024-12-01-at-8.38.46pm.png" alt=""></p>
        <blockquote>
        <p><a href="https://x.com/bsilva96/status/1861136096689606708">https://x.com/bsilva96/status/1861136096689606708</a></p>
        </blockquote>
        <p><a href="https://bsky.app/profile/k0kubun.com">@k0kubun</a> came through with insights into why things were slow and ways of improving the performance:</p>
        <p><img src="https://cdn.uploads.micro.blog/98548/2024/screenshot-2024-12-01-at-8.41.03pm.png" alt=""></p>
        <blockquote>
        <p><a href="https://x.com/k0kubun/status/1861149512640979260">https://x.com/k0kubun/status/1861149512640979260</a></p>
        </blockquote>
        <p>Let’s unpack his response. There are three parts to it:</p>
        <ol>
        <li><code>Range#each</code> is still written in C as of Ruby 3.4</li>
        <li><code>Integer#times</code> was converted from C to Ruby in Ruby 3.3</li>
        <li><code>Array#each</code> was converted from C to Ruby in Ruby 3.4</li>
        </ol>
        <h3 id="1-rangeeach-is-still-written-in-c-which-yjit-cant-optimize">1. <code>Range#each</code> is still written in C, which YJIT can’t optimize</h3>
        <p>Looking back at our Ruby code:</p>
        <pre><code>(0...10_000).each do |i|                     
          (0...100_000).each do |j|               
            a[i] += j % u                     
          end
          a[i] += r                      
        end
        </code></pre>
        <p>It’s written as a range, and range has its own <code>each</code> implementation, which is apparently written in C. The CRuby codebase is pretty easy to navigate - let’s find that implementation 🕵️‍♂️.</p>
        <p>Most core classes in Ruby have top-level C files named after them - in this case we’ve got <code>range.c</code> at the root of the project. CRuby has a pretty readable interface for exposing C functions as classes and methods - there is an <code>Init</code> function, usually at the bottom of the file. Inside that <code>Init</code> our classes, modules and methods are exposed from C to Ruby. Here are the relevant pieces of <code>Init_Range</code>:</p>
        <pre><code>void
        Init_Range(void)
        {
          //...
          rb_cRange = rb_struct_define_without_accessor(
            "Range", rb_cObject, range_alloc,
            "begin", "end", "excl", NULL);

          rb_include_module(rb_cRange, rb_mEnumerable);
          // ...
          rb_define_method(rb_cRange, "each", range_each, 0);
        </code></pre>
        <p>First, we define our <code>Range</code> class using <code>rb_struct_define...</code>. We name it <code>“Range”</code>, with a super class of <code>Object</code> (<code>rb_cObject</code>), and some initialization parameters (<code>“begin”</code>, <code>“end”</code> and whether to exclude the last value, ie the <code>..</code> vs <code>...</code> range syntax).</p>
        <p>Second, we include <code>Enumerable</code> using <code>rb_include_module</code>. That gives us all the great Ruby enumeration methods like <code>map</code>, <code>select</code>, <code>include?</code> and <a href="https://docs.ruby-lang.org/en/3.3/Enumerable.html">a bajillion others</a>. All you have to do is provide an <code>each</code> implementation and it handles the rest.</p>
        <p>Third, we define our <code>“each”</code> method. It’s implemented by the <code>range_each</code> function in C, and takes zero explicit arguments (blocks are not considered in this count).</p>
        <p><code>range_each</code> is hefty. It’s almost 100 lines long, and specializes into several versions of itself. I’ll highlight a few, collapsed all together:</p>
        <pre><code>static VALUE
        range_each(VALUE range)
        {
          //...
          range_each_fixnum_endless(beg);
          range_each_fixnum_loop(beg, end, range);
          range_each_bignum_endless(beg);
          rb_str_upto_endless_each(beg, sym_each_i, 0);
          // and even more...
        </code></pre>
        <p>These C functions handle all the variations of ranges you might use in your own code:</p>
        <div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-ruby" data-lang="ruby">(<span style="color:#ae81ff">0</span><span style="color:#f92672">...</span>)<span style="color:#f92672">.</span>each
        (<span style="color:#ae81ff">0</span><span style="color:#f92672">...</span><span style="color:#ae81ff">100</span>)<span style="color:#f92672">.</span>each
        (<span style="color:#e6db74">"a"</span><span style="color:#f92672">...</span><span style="color:#e6db74">"z"</span>)<span style="color:#f92672">.</span>each
        <span style="color:#75715e"># and on...</span>
        </code></pre></div><p>Why does it matter that <code>Range#each</code> is written in C? It means YJIT can’t inspect it - optimizations stop at the function call and resume when the function call returns. C functions are fast, but YJIT can take things further by creating specializations for hot paths of code. There is a great article from Aaron Patterson called <a href="https://railsatscale.com/2023-08-29-ruby-outperforms-c/">Ruby Outperforms C</a> where you can learn more about some of those specialized optimizations.</p>
        <h3 id="2-optimizing-our-loop-integertimes-was-converted-from-c-to-ruby-in-ruby-33">2. Optimizing our loop: <code>Integer#times</code> was converted from C to Ruby in Ruby 3.3</h3>
        <p>The hot path (<em>where most of our CPU time is spent</em>) is <code>Range#each</code>, which is a C function. YJIT can’t optimize C functions - they’re a black box. So what can we do?</p>
        <blockquote>
        <p>We converted Integer#times to Ruby in 3.3</p>
        </blockquote>
        <p>Interesting! In Ruby 3.3, <code>Integer#times</code> was <a href="https://github.com/ruby/ruby/pull/8388">converted from a C function to a Ruby method</a>! Here’s the 3.3+ version - its pretty simple:</p>
        <pre><code>def times
          #... a little C interop code
          i = 0
          while i &lt; self
            yield i
            i = i.succ
          end
          self
        end
        </code></pre>
        <p>Very simple. It’s just a basic while loop. Most importantly, it’s all Ruby code, which means YJIT should be able to introspect and optimize it!</p>
        <h3 id="an-aside-on-integersucc">An aside on <code>Integer#succ</code></h3>
        <p>The slightly odd part of that code is <code>i.succ</code>. I’d never heard of <code>Integer#succ</code>, which apparently gives you the “successor” to an integer.</p>
        <p><img src="https://cdn.uploads.micro.blog/98548/2024/0c8bd56f64.png" alt=""></p>
        <blockquote>
        <p>I’ve never seen this show, and yet it’s the first thing I thought of when I learned about this method. Thanks, advertising.</p>
        </blockquote>
        <p>There was a PR to improve the performance of <code>Integer#succ</code> in early 2024, which helped me understand why anyone would ever use it:</p>
        <blockquote>
        <p>We use Integer#succ when we rewrite loop methods in Ruby (e.g. Integer#times and Array#each) because opt_succ (i = i.succ) is faster to dispatch on the interpreter than putobject 1; opt_plus (i += 1).</p>
        <p><a href="https://github.com/ruby/ruby/pull/9519">https://github.com/ruby/ruby/pull/9519</a></p>
        </blockquote>
        <p><code>Integer#succ</code> is like a virtual machine cheat code. It takes a common operation (adding 1 to an integer) and turns it from two virtual machine operations into one. We can call <code>disasm</code> on the <code>times</code> method to see that in action:</p>
        <pre><code>puts RubyVM::InstructionSequence.disasm(1.method(:times))
        </code></pre>
        <p>The <code>Integer#times</code> method gets broken down into a lot of Ruby VM bytecode, but we only care about a few lines:</p>
        <pre><code>...
        0025 getlocal_WC_0   i@0
        0027 opt_succ        &lt;calldata!mid:succ, ARGS_SIMPLE&gt;[CcCr]
        0029 setlocal_WC_0   i@0
        ...
        </code></pre>
        <ul>
        <li><code>getlocal_WC_0</code> gets our <code>i</code> variable from the current scope. That’s the <code>i</code> in <code>i.succ</code></li>
        <li><code>opt_succ</code> performs the <code>succ</code> call in our <code>i.succ</code>. It will either call the actual <code>Integer#succ</code> method, or an optimized C function for small numbers</li>
        <li>In Ruby 3.4 with YJIT enabled, small numbers get optimized even further into machine code (just a note, not shown in the VM machine code)</li>
        <li><code>setlocal_WC_0</code> sets the result of <code>opt_succ</code> to our local variable <code>i</code></li>
        </ul>
        <p>If we change from <code>i = i.succ</code> to <code>i += 1</code>, we now have two VM operations take the place of <code>opt_succ</code>:</p>
        <pre><code>...
        0025 getlocal_WC_0        i@0
        0027 putobject_INT2FIX_1_
        0028 opt_plus             &lt;calldata!mid:+, argc:1, ARGS_SIMPLE&gt;
        0029 setlocal_WC_0        i@0
        ...
        </code></pre>
        <p>Everything is essentially the same as before, except now we have two steps to go through instead of one:</p>
        <ul>
        <li><code>putobject_INT2FIX_1_</code> pushes the integer <code>1</code> onto the virtual machine stack</li>
        <li><code>opt_plus</code> is the <code>+</code> in our <code>+= 1</code>, and calls either the Ruby <code>+</code> method or an optimized C function for small numbers</li>
        <li>There is probably a YJIT optimization for <code>opt_plus</code> as well</li>
        </ul>
        <p>If there is nothing else to learn from this code, it’s this: the kinds of optimizations you do at the VM and JIT level are <em>deep</em>. When writing general Ruby programs we typically don’t and <em>shouldn’t</em> consider the impact of one versus two <em>machine code instructions</em>. But at the JIT level, on the scale of millions and billions of operations, it matters!</p>
        <h3 id="back-to-integertimes">Back to <code>Integer#times</code></h3>
        <p>Let’s try running our benchmark code again, using <code>times</code>! Instead of iterating over ranges, we simply iterate for <code>10_000</code> and <code>100_000</code> <code>times</code>:</p>
        <div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-ruby" data-lang="ruby">u <span style="color:#f92672">=</span> <span style="color:#66d9ef">ARGV</span><span style="color:#f92672">[</span><span style="color:#ae81ff">0</span><span style="color:#f92672">].</span>to_i        
        r <span style="color:#f92672">=</span> rand(<span style="color:#ae81ff">10_000</span>)        
        a <span style="color:#f92672">=</span> Array<span style="color:#f92672">.</span>new(<span style="color:#ae81ff">10_000</span>, <span style="color:#ae81ff">0</span>)
        	
        <span style="color:#ae81ff">10_000</span><span style="color:#f92672">.</span>times <span style="color:#66d9ef">do</span> <span style="color:#f92672">|</span>i<span style="color:#f92672">|</span>
          <span style="color:#ae81ff">100_000</span><span style="color:#f92672">.</span>times <span style="color:#66d9ef">do</span> <span style="color:#f92672">|</span>j<span style="color:#f92672">|</span>
            a<span style="color:#f92672">[</span>i<span style="color:#f92672">]</span> <span style="color:#f92672">+=</span> j <span style="color:#f92672">%</span> u
          <span style="color:#66d9ef">end</span>
          a<span style="color:#f92672">[</span>i<span style="color:#f92672">]</span> <span style="color:#f92672">+=</span> r
        <span style="color:#66d9ef">end</span>
        	
        puts a<span style="color:#f92672">[</span>r<span style="color:#f92672">]</span>
        </code></pre></div><table>
        <thead>
        <tr>
        <th></th>
        <th>Loops</th>
        </tr>
        </thead>
        <tbody>
        <tr>
        <td>Range#each</td>
        <td>25.57s</td>
        </tr>
        <tr>
        <td>Integer#times</td>
        <td>13.66s</td>
        </tr>
        </tbody>
        </table>
        <p>Nice! YJIT makes a much larger impact using <code>Integer#times</code>. That trims things down significantly, taking it down to 13.66 seconds on my machine. On <a href="https://bsky.app/profile/k0kubun.com">@k0kobun</a>’s machine it actually goes down to 9 seconds (and 8 seconds on Ruby 3.4).</p>
        <blockquote>
        <p>It’s probably Ruby 3.5’s job to make it faster than 8s though.</p>
        </blockquote>
        <p>We might look forward to even faster performance in Ruby 3.5. We’ll see!</p>
        <h3 id="3-arrayeach-was-converted-from-c-to-ruby-in-ruby-34">3. <code>Array#each</code> was converted from C to Ruby in Ruby 3.4</h3>
        <p>CRuby continues to see C code rewritten in Ruby, and in Ruby 3.4 <code>Array#each</code> was one of those changes. Here is an <a href="https://github.com/ruby/ruby/pull/6687/files">example of the first attempt at implementing it</a>:</p>
        <pre><code>def each
          unless block_given?
            return to_enum(:each) { self.length }
          end
          i = 0
          while i &lt; self.length
            yield self[i]
            i = i.succ
          end
          self
        end
        </code></pre>
        <p>Super simple and readable! And YJIT optimizable!</p>
        <p>Unfortunately, due to something related to CRuby internals, it contained <a href="https://jpcamara.com/2024/06/23/your-ruby-programs.html#race-conditions">race conditions</a>. A later implementation <a href="https://github.com/ruby/ruby/pull/11955">landed in Ruby 3.4</a>.</p>
        <pre><code>def each
          Primitive.attr! :inline_block, :c_trace

          unless defined?(yield)
            return Primitive.cexpr! 'SIZED_ENUMERATOR(self, 0, 0, ary_enum_length)'
          end
          _i = 0
          value = nil
          while Primitive.cexpr!(%q{ ary_fetch_next(self, LOCAL_PTR(_i), LOCAL_PTR(value)) })
            yield value
          end
          self
        end
        </code></pre>
        <p>Unlike the first implementation, and unlike <code>Integer#times</code>, things are a bit more cryptic this time. This is definitely not pure Ruby code that anyone could be expected to write. Somehow, the <code>Primitive</code> module seems to allow evaluating C code from Ruby, and in doing so avoids the race conditions present in the pure Ruby solution<sup id="fnref:3"><a href="#fn:3" class="footnote-ref" role="doc-noteref">3</a></sup>.</p>
        <p>By fetching indexes and values using C code, I think it results in a more atomic operation. I have no idea why the <code>Primitive.cexpr!</code> is used to return the enumerator, or what value <code>Primitive.attr! :inline_block</code> provides. Please comment if you have insights there!</p>
        <p>I was a little loose with my earlier <code>Integer#times</code> source code as well. That actually had a bit of this <code>Primitive</code> syntax as well. The core of the method is what we looked at, and it’s all Ruby, but the start of the method contains the same <code>Primitive</code> calls for <code>:inline_block</code> and returning the enumerator:</p>
        <pre><code>def times
          Primitive.attr! :inline_block
          unless defined?(yield)
            return Primitive.cexpr! 'SIZED_ENUMERATOR(self, 0, 0, int_dotimes_size)'
          end
          #...
        </code></pre>
        <p>Ok - it’s more cryptic than <code>Integer#times</code> was, but <code>Array#each</code> is mostly Ruby (on Ruby 3.4+). Let’s give it a try using arrays instead of ranges or <code>times</code>:</p>
        <div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-ruby" data-lang="ruby">u <span style="color:#f92672">=</span> <span style="color:#66d9ef">ARGV</span><span style="color:#f92672">[</span><span style="color:#ae81ff">0</span><span style="color:#f92672">].</span>to_i
        r <span style="color:#f92672">=</span> rand(<span style="color:#ae81ff">10_000</span>)
        a <span style="color:#f92672">=</span> Array<span style="color:#f92672">.</span>new(<span style="color:#ae81ff">10_000</span>, <span style="color:#ae81ff">0</span>)
        	
        outer <span style="color:#f92672">=</span> (<span style="color:#ae81ff">0</span><span style="color:#f92672">...</span><span style="color:#ae81ff">10_000</span>)<span style="color:#f92672">.</span>to_a<span style="color:#f92672">.</span>freeze
        inner <span style="color:#f92672">=</span> (<span style="color:#ae81ff">0</span><span style="color:#f92672">...</span><span style="color:#ae81ff">100_000</span>)<span style="color:#f92672">.</span>to_a<span style="color:#f92672">.</span>freeze
        outer<span style="color:#f92672">.</span>each <span style="color:#66d9ef">do</span> <span style="color:#f92672">|</span>i<span style="color:#f92672">|</span>
          inner<span style="color:#f92672">.</span>each <span style="color:#66d9ef">do</span> <span style="color:#f92672">|</span>j<span style="color:#f92672">|</span>
            a<span style="color:#f92672">[</span>i<span style="color:#f92672">]</span> <span style="color:#f92672">+=</span> j <span style="color:#f92672">%</span> u
          <span style="color:#66d9ef">end</span>
          a<span style="color:#f92672">[</span>i<span style="color:#f92672">]</span> <span style="color:#f92672">+=</span> r
        <span style="color:#66d9ef">end</span>
        	
        puts a<span style="color:#f92672">[</span>r<span style="color:#f92672">]</span>
        </code></pre></div><p>Despite the embedded C code, YJIT still seems capable of making some hefty performance optimizations. It’s within the same range as <code>Integer#times</code>!</p>
        <table>
        <thead>
        <tr>
        <th></th>
        <th>Loops</th>
        </tr>
        </thead>
        <tbody>
        <tr>
        <td>Range#each</td>
        <td>25.57s</td>
        </tr>
        <tr>
        <td>Integer#times</td>
        <td>13.66s</td>
        </tr>
        <tr>
        <td>Array#each</td>
        <td>13.96s</td>
        </tr>
        </tbody>
        </table>
        <h3 id="microbenchmarking-ruby-performance">Microbenchmarking Ruby performance</h3>
        <p>I’ve forked the original language implementation repo, and created my own repository called “Ruby Microbench”. It takes all of the examples discussed, as well as several other forms of doing the iteration in Ruby: <a href="https://github.com/jpcamara/ruby_microbench">https://github.com/jpcamara/ruby_microbench</a></p>
        <p>Here is the output of just running those using Ruby 3.4 with and without YJIT:</p>
        <table>
        <thead>
        <tr>
        <th></th>
        <th>fibonacci</th><th>array#each</th>
        <th>range#each</th>
        <th>times</th>
        <th>for</th>
        <th>while</th>
        <th>loop do</th>
        </tr>
        </thead>
        <tbody>
        <tr>
        <td>Ruby 3.4 YJIT</td>
        <td>2.19s</td>
        <td>14.02s</td>
        <td>26.61s</td>
        <td>13.12s</td>
        <td>14.91s</td>
        <td>37.10s</td>
        <td>13.95s</td></tr>
        <tr>
        <td>Ruby 3.4</td>
        <td>16.49s</td>
        <td>34.29s</td>
        <td>33.88s</td>
        <td>33.18s</td>
        <td>36.32s</td>
        <td>37.14s</td>
        <td>50.65s</td>
        </tr>
        </tbody>
        </table>
        <p>I have no idea why the <code>while</code> loop example I wrote seems to be so slow. I’d expected it to run much faster. Maybe there’s an issue with how I wrote it - feel free to open an issue or PR if you see something wrong with my implementation. The <code>loop do</code> (taken from <a href="https://bsky.app/profile/timtilberg.bsky.social">@timtilberg</a>’s <a href="https://x.com/timtilberg/status/1861194052516864004">example</a>) runs around the same speed as <code>Integer#times</code> - although its performance is <em>awful</em> with YJIT turned off.</p>
        <blockquote>
        <p>📝 The <code>for in</code> loop and <code>array#each</code> have very similar performance, and that’s because at the Ruby VM bytecode level they are almost identical. <code>for in</code> is mostly syntactic sugar that transforms into an <code>#each</code> call in the VM. Thanks to <a href="https://ruby.social/@dodecadaniel">Daniel Colson</a> for pointing this out, and you can read his <a href="https://danieljamescolson.com/blog/for-loops">for loops in Ruby</a> article for some additional information and nuance around <code>for in</code>!</p>
        </blockquote>
        <p>In addition to running Ruby 3.4, for fun I have it using <code>rbenv</code> to run:</p>
        <ul>
        <li>Ruby 3.3</li>
        <li>Ruby 3.3 YJIT</li>
        <li>Ruby 3.2</li>
        <li>Ruby 3.2 YJIT</li>
        <li>TruffleRuby 24.1</li>
        <li>Ruby Artichoke</li>
        <li>MRuby</li>
        </ul>
        <p>A few of the test runs are listed here:</p>
        <table>
        <thead>
        <tr>
        <th></th>
        <th>fibonacci</th>
        <th>
        array#each
        </th>
        <th>range#each</th>
        <th>times</th>
        <th>for</th>
        <th>while</th>
        <th>loop do</th>
        </tr>
        </thead>
        <tbody>
        <tr>
        <td>Ruby 3.4 YJIT</td><td>2.19s</td>
        <td>14.02s</td>
        <td>26.61s</td>
        <td>13.12s</td>
        <td>14.91s</td>
        <td>37.10s</td>
        <td>13.95s</td></tr>
        <tr>
        <td>Ruby 3.4</td>
        <td>16.49s</td>
        <td>34.29s</td><td>33.88s</td>
        <td>33.18s</td>
        <td>36.32s</td>
        <td>37.14s</td>
        <td>50.65s</td>
        </tr>
        <tr><td>TruffleRuby 24.1</td><td>0.92s</td>
        <td>0.97s</td>
        <td>0.92s</td><td>2.39s</td>
        <td>2.06s</td><td>3.90s</td>
        <td>0.77s</td>
        </tr><tr>
        <td>MRuby 3.3</td>
        <td>28.83s</td>
        <td>144.65s</td>
        <td>126.40s</td>
        <td>128.22s</td>
        <td>133.58s</td><td>91.55s</td>
        <td>144.93s</td>
        </tr>
        <tr>
        <td>Artichoke</td>
        <td>19.71s</td>
        <td>236.10s</td><td>214.55s</td>
        <td>214.51s</td>
        <td>215.95s</td>
        <td>174.70s</td>
        <td>264.67s</td>
        </tr>
        </tbody>
        </table>
        <p>Based on that, I’ve taken the original visualization and made a Ruby specific one here just for the <code>fibonacci</code> run:</p>
        <div id="ruby_visual"></div>
        <h3 id="speeding-up-rangeeach">Speeding up <code>range#each</code></h3>
        <p>Can we, the non <a href="https://bsky.app/profile/k0kubun.com">@k0kobun</a>’s of the world, make <code>range#each</code> faster? If I monkey patch the <code>Range</code> class with a pure-ruby implementation, things <em>do</em> get much faster! Here’s my implementation:</p>
        <pre><code>class Range
          def each
            beginning = self.begin
            ending = self.end
            i = beginning
            loop do
              break if i == ending
              yield i
              i = i.succ
            end
          end
        end
        </code></pre>
        <p>And here is the change in performance - 2 seconds slower than <code>times</code> - not bad!</p>
        <table>
        <thead>
        <tr>
        <th>				
        </th>
        <th>
        Time spent
        </th>
        </tr>
        </thead>
        <tbody>
        <tr>
        <td>Range#each in C</td>
        <td>25.57s</td>
        </tr>
        <tr>
        <td>Range#each in Ruby</td>
        <td>16.64s</td>
        </tr>
        </tbody>
        </table>
        <p>This is obviously over-simplified. I don’t handle all of the different cases of <code>Range</code>, and there may be nuances I am missing. Also, most of the Ruby rewritten methods I’ve seen invoke a <code>Primitive</code> class for certain operations. I’d love to learn more about when and why it’s needed.</p>
        <p>But! It goes to show the power of moving things <em>out</em> of C and letting YJIT optimize our code. It can improve performance in ways that would be difficult or impossible to replicate in regular C code.</p>
        <h3 id="yjit-standard-library">YJIT standard library</h3>
        <p>Last year Aaron Patterson wrote an article called <a href="https://railsatscale.com/2023-08-29-ruby-outperforms-c/">Ruby Outperforms C</a>, in which he rewrote a C extension in Ruby for some GraphQL parsing. The Ruby code outperformed C thanks to YJIT optimizations.</p>
        <p>This got me thinking that it would be interesting to see a kind of “YJIT standard library” emerge, where core ruby functionality run in C could be swapped out for Ruby implementations for use by people using YJIT.</p>
        <p>As it turns out, this is almost exactly what the core YJIT team has been doing. In many cases they’ve completely removed C code, but more recently they’ve created a <code>with_yjit</code> block. The code will only take effect if YJIT is enabled, and otherwise the C code will run. For example, this is how<code>Array#each</code> is implemented:</p>
        <pre><code>with_yjit do
          if Primitive.rb_builtin_basic_definition_p(:each)
            undef :each

            def each # :nodoc:
              # ... we examined this code earlier ...
            end
          end
        end
        </code></pre>
        <p>As of Ruby 3.3, YJIT can be lazily initialized. Thankfully the <code>with_yjit</code> code handles this - the appropriate <code>with_yjit</code> versions of methods will be run once YJIT is enabled:</p>
        <pre><code># Uses C-builtin
        [1, 2, 3].each do |i|
          puts i
        end

        RubyVM::YJIT.enable

        # Uses Ruby version, which can be YJIT optimized
        [1, 2, 3].each do |i|
          puts i
        end
        </code></pre>
        <p>This is because <code>with_yjit</code> is a YJIT “hook”, which is called the moment YJIT is enabled. After being called, it is removed from the runtime using <code>undef :with_yjit</code>.</p>
        <h3 id="investigating-yjit-optimizations">Investigating YJIT optimizations</h3>
        <p>We’ve looked at Ruby code. We’ve looked at C code. We’ve looked at Ruby VM bytecode. Why not take it one step deeper and look at some <em>machine code</em>? And maybe some Rust code? Hey - where are you going! Don’t walk away while I’m talking to you!</p>
        <p>If you <em>haven’t</em> walked away, or skipped to the next section, let’s take a look at a small sliver of YJIT while we’re here!</p>
        <p>We can see the machine code YJIT generates 😱. It’s possible by building CRuby from source with YJIT debug flags. If you’re on a Mac you can see <a href="https://jpcamara.com/2024/12/02/my-macos-setup.html">my MacOS setup for hacking on CRuby</a> or <a href="https://jpcamara.com/2024/11/27/my-docker-setup.html">my docker setup for hacking on CRuby</a> for more elaborate instructions on building Ruby. But the simplified step is when you go to <code>./configure</code> Ruby, you hand in an option of <code>--enable-yjit=dev</code>:</p>
        <pre><code>./configure --enable-yjit=dev
        make install
        </code></pre>
        <p>Let’s use our <code>Integer#times</code> example from earlier as our example Ruby code:</p>
        <div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-ruby" data-lang="ruby">u <span style="color:#f92672">=</span> <span style="color:#66d9ef">ARGV</span><span style="color:#f92672">[</span><span style="color:#ae81ff">0</span><span style="color:#f92672">].</span>to_i
        r <span style="color:#f92672">=</span> rand(<span style="color:#ae81ff">10_000</span>)
        a <span style="color:#f92672">=</span> Array<span style="color:#f92672">.</span>new(<span style="color:#ae81ff">10_000</span>, <span style="color:#ae81ff">0</span>)
        	
        <span style="color:#ae81ff">10_000</span><span style="color:#f92672">.</span>times <span style="color:#66d9ef">do</span> <span style="color:#f92672">|</span>i<span style="color:#f92672">|</span>
          <span style="color:#ae81ff">100_000</span><span style="color:#f92672">.</span>times <span style="color:#66d9ef">do</span> <span style="color:#f92672">|</span>j<span style="color:#f92672">|</span>
            a<span style="color:#f92672">[</span>i<span style="color:#f92672">]</span> <span style="color:#f92672">+=</span> j <span style="color:#f92672">%</span> u
          <span style="color:#66d9ef">end</span>
          a<span style="color:#f92672">[</span>i<span style="color:#f92672">]</span> <span style="color:#f92672">+=</span> r
        <span style="color:#66d9ef">end</span>
        	
        puts a<span style="color:#f92672">[</span>r<span style="color:#f92672">]</span>
        </code></pre></div><p>Because you’ve built Ruby with YJIT in dev mode, you can hand in the <code>--yjit-dump-disasm</code> flag when running your ruby program:</p>
        <pre><code>./ruby --yjit --yjit-dump-disasm test.rb 40
        </code></pre>
        <p>Using this, we can see the machine code created. We’ll just focus in on one tiny part - the machine code equivalent of the Ruby VM bytecode we read earlier. Here is the original VM bytecode for <code>opt_succ</code>, which is generated when you call <code>i.succ</code>, the <code>Integer#succ</code> method:</p>
        <pre><code>...
        0027 opt_succ        &lt;calldata!mid:succ, ARGS_SIMPLE&gt;[CcCr]
        ...
        </code></pre>
        <p>And here is the machine code YJIT generates in this scenario, on my Mac M2 arm64 architecture:</p>
        <pre><code># Block: times@&lt;internal:numeric&gt;:259 
        # reg_mapping: [Some(Stack(0)), None, None, None, None]
        # Insn: 0027 opt_succ (stack_size: 1)
        # call to Integer#succ
        # guard object is fixnum
        0x1096808c4: tst x1, #1
        0x1096808c8: b.eq #0x109683014
        0x1096808cc: nop 
        0x1096808d0: nop 
        0x1096808d4: nop 
        0x1096808d8: nop 
        0x1096808dc: nop 
        # Integer#succ
        0x1096808e0: adds x11, x1, #2
        0x1096808e4: b.vs #0x109683048
        0x1096808e8: mov x1, x11
        </code></pre>
        <p>To be honest, I about 25% understand this, and 75% am combining my own logic and AI to learn it 🤫. Feel free to yell at me if I get it a little wrong, I’d love to learn more. But here’s how I break this down.</p>
        <pre><code># Block: times@&lt;internal:numeric&gt;:259
        </code></pre>
        <p>👆🏼This roughly corresponds to the line <code>i = i.succ</code> in the <code>Integer#times</code> method in <code>numeric.rb</code>. I say roughly because in my current code I see that on line 258, but maybe it shows the end of the block it’s run in since YJIT compiles “blocks” of code:</p>
        <pre><code>256: while i &lt; self
        257:   yield i
        258:   i = i.succ
        259: end

        # reg_mapping: [Some(Stack(0)), None, None, None, None]
        # Insn: 0027 opt_succ (stack_size: 1)
        # call to Integer#succ
        </code></pre>
        <p>👆🏼I have no idea what <code>reg_mapping</code> means - probably mapping how it uses a CPU register? <code>Insn: 0027 opt_succ</code> looks very familiar! That’s our VM bytecode! <code>call to Integer#succ</code> is just a helpful comment added. YJIT is capable of adding comments to the machine code. We still haven’t even left the safety of the comments 😅.</p>
        <pre><code># guard object is fixnum
        </code></pre>
        <p>👆🏼This is interesting. I can find a corresponding bit of Rust code that maps directly to this. Let’s take a look at it:</p>
        <pre><code>fn jit_rb_int_succ(
          //...
          asm: &amp;mut Assembler,
          //...
        ) -&gt; bool {
          // Guard the receiver is fixnum
          let recv_type = asm.ctx.get_opnd_type(StackOpnd(0));
          let recv = asm.stack_pop(1);
          if recv_type != Type::Fixnum {
            asm_comment!(asm, "guard object is fixnum");
            asm.test(recv, Opnd::Imm(RUBY_FIXNUM_FLAG as i64));
                 asm.jz(Target::side_exit(Counter::opt_succ_not_fixnum));
          }

          asm_comment!(asm, "Integer#succ");
          let out_val = asm.add(recv, Opnd::Imm(2)); // 2 is untagged Fixnum 1
          asm.jo(Target::side_exit(Counter::opt_succ_overflow));

          // Push the output onto the stack
          let dst = asm.stack_push(Type::Fixnum);
          asm.mov(dst, out_val);

          true
        }
        </code></pre>
        <p>Oh nice! This is the actual YJIT Rust implementation of the <code>opt_succ</code> call. This is that optimization <a href="https://bsky.app/profile/k0kubun.com">@k0kobun</a> made to further improve <code>opt_succ</code> performance beyond the bytecode C function calls. We’re in the section that is checking if what we’re operating on is a Fixnum, which is a way small integers are stored internally in CRuby:</p>
        <pre><code>if recv_type != TypeFixnum 
          asm_comment!(asm, "guard object is fixnum");
          asm.test(recv, Opnd::Imm(RUBY_FIXNUM_FLAG as i64));
          asm.jz(Target::side_exit(Counter::opt_succ_not_fixnum));
        }
        </code></pre>
        <p>That becomes this machine code:</p>
        <pre><code># guard object is fixnum
        0x1096808c4: tst x1, #1
        0x1096808c8: b.eq #0x109683014
        </code></pre>
        <p><code>asm.test</code> generates <code>tst x1, #1</code>, which according to an AI bot I asked is checking the least significant bit, which is a Fixnum “tag” that indicates this is a Fixnum. If it’s Fixnum, the result is 1 and <code>b.eq</code> is false. If it’s not a Fixnum, the result is <code>0</code> and <code>b.eq</code> is true and jumps away from this code.</p>
        <pre><code>0x1096808cc: nop 
        0x1096808d0: nop 
        0x1096808d4: nop 
        0x1096808d8: nop 
        0x1096808dc: nop 
        </code></pre>
        <p>🤖 “NOPs for alignment/padding”. Thanks AI. I don’t know why it is needed, but at least I know what it probably is.</p>
        <p>Finally, we <em>actually</em> add 1 to the number.</p>
        <pre><code>asm_comment!(asm, "Integer#succ");
        let out_val = asm.add(recv, Opnd::Imm(2)); // 2 is untagged Fixnum 1
        asm.jo(Target::side_exit(Counter::opt_succ_overflow));

        // Push the output onto the stack
        let dst = asm.stack_push(Type::Fixnum);
        asm.mov(dst, out_val);
        </code></pre>
        <p>The Rust code generates our <code>Integer#succ</code> comment. Then, to add 1, because of the “Fixnum tag” data embedded within our integer, actually means we have to add 2 using <code>adds x11, x1, #2</code> 😵‍💫. If we overflow the space available, it exits to a different code path - <code>b.vs</code> is a branch on overflow. Otherwise, it stores the result with <code>mov x1, x11</code>!</p>
        <pre><code># Integer#succ
        0x1096808e0: adds x11, x1, #2
        0x1096808e4: b.vs #0x109683048
        0x1096808e8: mov x1, x11
        </code></pre>
        <p>😮‍💨. That was a lot. And it seems like <em>alot</em> of working is being done, but because it’s such low level machine code it’s presumably super fast. We examined a teensy tiny portion of what YJIT is capable of generating - JITs are complicated!</p>
        <p>Thanks to <a href="https://bsky.app/profile/k0kubun.com">@k0kobun</a> for providing me with the commands and pointing me at the <a href="https://github.com/ruby/ruby/blob/master/doc/yjit/yjit.md">YJIT docs</a> which contain tons of additional options as well.</p>
        <h3 id="the-future-of-cruby-optimizations">The future of CRuby optimizations</h3>
        <p>The irony of language implementation is that you often work less in the language you’re implementing than you do in something lower-level - in Ruby’s case, that’s mostly C and some Rust.</p>
        <p>With a layer like YJIT, it potentially opens up a future where more of the language becomes plain Ruby, and Ruby developer contribution is easier. Many languages have a smaller low level core, and the majority of the language is written in itself (like Java, for instance). Maybe that’s a future for CRuby, someday! Until then, keep the YJIT optimizations coming, YJIT team!</p>
        <script src="https://jpcamara.com/d3.v7.min.js"></script>
        <script src="https://jpcamara.com/latency_visual.js"></script>
        <section class="footnotes" role="doc-endnotes">
        <hr>
        <ol>
        <li id="fn:1" role="doc-endnote">
        <p>Naive in this case meaning that there are more efficient ways to calculate fibonacci numbers in a program&nbsp;<a href="#fnref:1" class="footnote-backref" role="doc-backlink">↩︎</a></p>
        </li>
        <li id="fn:2" role="doc-endnote">
        <p>MJIT, the precursor to YJIT, made Ruby much faster on certain benchmarks. But on large realistic Rails applications it actually <a href="https://bugs.ruby-lang.org/issues/14490">made things <em>slower</em></a>&nbsp;<a href="#fnref:2" class="footnote-backref" role="doc-backlink">↩︎</a></p>
        </li>
        <li id="fn:3" role="doc-endnote">
        <p>When C code is running, it has to opt-in to releasing the GVL, so it’s more difficult for threads to corrupt or modify data mid-operation. The original Ruby version could yield the GVL at points that would invalidate the array. That’s my understanding of the situation anyways.&nbsp;<a href="#fnref:3" class="footnote-backref" role="doc-backlink">↩︎</a></p>
        </li>
        </ol>
        </section>
      text: |-
        There is a recent language comparison repo which has been getting shared a lot. In it, CRuby was the third slowest option, only beating out R and Python.
        The repo author, @BenjDicken, created a fun visualization of each language’s performance. Here’s one of the visualizations, which shows Ruby as the third slowest language benchmarked:


        The code for this visualization is from https://benjdd.com/languages/, with permission from @BenjDicken

        The repository describes itself as:

        A repo for collaboratively building small benchmarks to compare languages.

        It contains two different benchmarks:

        “Loops”, which “Emphasizes loop, conditional, and basic math performance”
        “Fibonacci”, which “Emphasizes function call overhead and recursion.”

        The loop example iterates 1 billion times, utilizing a nested loop:
        u = ARGV[0].to_i       
        r = rand(10_000)                          
        a = Array.new(10_000, 0)                 
        	
        (0...10_000).each do |i|                     
          (0...100_000).each do |j|               
            a[i] += j % u                     
          end
          a[i] += r                      
        end
        	
        puts a[r]
        The Fibonacci example is a basic “naive” Fibonacci implementation1:
        def fibonacci(n)
          return 0 if n == 0
          return 1 if n == 1
          fibonacci(n - 1) + fibonacci(n - 2)
        end

        u = ARGV[0].to_i
        r = 0

        (1...u).each do |i|
          r += fibonacci(i)
        end

        puts r

        Run on @BenjDicken’s M3 MacBook Pro, Ruby 3.3.6 takes 28 seconds to run the loop iteration example, and 12 seconds to run the Fibonacci example. For comparison, node.js takes a little over a second for both examples - it’s not a great showing for Ruby.

          
            
              
              Fibonacci
              Loops
            
          
          
            
              Ruby
              12.17s
              28.80s
            
            
              node.js
              1.11s
              1.03s
            
          

        From this point on, I’ll use benchmarks relative to my own computer. Running the same benchmark on my M2 MacBook Air, I get 33.43 seconds for the loops and 16.33 seconds for fibonacci - even worse 🥺. Node runs a little over 1 second for fibonacci and 2 seconds for the loop example.

          
            
              
              Fibonacci
              Loops
            
          
          
            
              Ruby
              16.33s
              33.43s
            
            
              node.js
              1.36s
              2.07s
            
          

        Who cares?
        In most ways, these types of benchmarks are meaningless. Python was the slowest language in the benchmark, and yet at the same time it’s the most used language on Github as of October 2024. Ruby runs some of the largest web apps in the world. I ran a benchmark recently of websocket performance between the Ruby Falcon web server and node.js, and the Ruby results were close to the node.js results. Are you doing a billion loop iterations or using web sockets?
        A programming language should be reasonably efficient - after that the usefulness of the language, the type of tasks you work on, and language productivity outweigh the speed at which you can run a billion iterations of a loop, or complete an intentionally inefficient implementation of a Fibonacci method.
        That said:

        The programming world loves microbenchmarks 🤷‍♂️
        Having a fast benchmark may not be valuable in practice but it has meaning for people’s interest in a language. Some would claim it means you’ll have an easier time scaling performance, but that’s arguable2
        It’s disappointing if your language of choice doesn’t perform well. It’s nice to be able to say “I use and enjoy this language, and it runs fast in all benchmarks!”

        In the case of this Ruby benchmark, I had a feeling that YJIT wasn’t being applied in the Ruby code, so I checked the repo. Lo and behold, the command was as follows:
        ruby ./code.rb 40

        We know my results from earlier (33 seconds and 16 seconds). What do we get with YJIT applied?
        ruby --yjit ./code.rb 40


          
            
              
        Fibonacci
        Loops




        Ruby
        2.06s
        25.57s



        Nice! With YJIT, Fibonacci gets a massive boost - going from 16.88 seconds down to 2.06 seconds. It’s close to the speed of node.js at that point!
        YJIT makes a more modest difference for the looping example - going from 33.43 seconds down to 25.57 seconds. Why is that?
        A team effort
        I wasn’t alone in trying out these code samples with YJIT. On twitter, @bsilva96 had asked the same questions:


        https://x.com/bsilva96/status/1861136096689606708

        @k0kubun came through with insights into why things were slow and ways of improving the performance:


        https://x.com/k0kubun/status/1861149512640979260

        Let’s unpack his response. There are three parts to it:

        Range#each is still written in C as of Ruby 3.4
        Integer#times was converted from C to Ruby in Ruby 3.3
        Array#each was converted from C to Ruby in Ruby 3.4

        1. Range#each is still written in C, which YJIT can’t optimize
        Looking back at our Ruby code:
        (0...10_000).each do |i|                     
          (0...100_000).each do |j|               
            a[i] += j % u                     
          end
          a[i] += r                      
        end

        It’s written as a range, and range has its own each implementation, which is apparently written in C. The CRuby codebase is pretty easy to navigate - let’s find that implementation 🕵️‍♂️.
        Most core classes in Ruby have top-level C files named after them - in this case we’ve got range.c at the root of the project. CRuby has a pretty readable interface for exposing C functions as classes and methods - there is an Init function, usually at the bottom of the file. Inside that Init our classes, modules and methods are exposed from C to Ruby. Here are the relevant pieces of Init_Range:
        void
        Init_Range(void)
        {
          //...
          rb_cRange = rb_struct_define_without_accessor(
            "Range", rb_cObject, range_alloc,
            "begin", "end", "excl", NULL);

          rb_include_module(rb_cRange, rb_mEnumerable);
          // ...
          rb_define_method(rb_cRange, "each", range_each, 0);

        First, we define our Range class using rb_struct_define.... We name it “Range”, with a super class of Object (rb_cObject), and some initialization parameters (“begin”, “end” and whether to exclude the last value, ie the .. vs ... range syntax).
        Second, we include Enumerable using rb_include_module. That gives us all the great Ruby enumeration methods like map, select, include? and a bajillion others. All you have to do is provide an each implementation and it handles the rest.
        Third, we define our “each” method. It’s implemented by the range_each function in C, and takes zero explicit arguments (blocks are not considered in this count).
        range_each is hefty. It’s almost 100 lines long, and specializes into several versions of itself. I’ll highlight a few, collapsed all together:
        static VALUE
        range_each(VALUE range)
        {
          //...
          range_each_fixnum_endless(beg);
          range_each_fixnum_loop(beg, end, range);
          range_each_bignum_endless(beg);
          rb_str_upto_endless_each(beg, sym_each_i, 0);
          // and even more...

        These C functions handle all the variations of ranges you might use in your own code:
        (0...).each
        (0...100).each
        ("a"..."z").each
        # and on...
        Why does it matter that Range#each is written in C? It means YJIT can’t inspect it - optimizations stop at the function call and resume when the function call returns. C functions are fast, but YJIT can take things further by creating specializations for hot paths of code. There is a great article from Aaron Patterson called Ruby Outperforms C where you can learn more about some of those specialized optimizations.
        2. Optimizing our loop: Integer#times was converted from C to Ruby in Ruby 3.3
        The hot path (where most of our CPU time is spent) is Range#each, which is a C function. YJIT can’t optimize C functions - they’re a black box. So what can we do?

        We converted Integer#times to Ruby in 3.3

        Interesting! In Ruby 3.3, Integer#times was converted from a C function to a Ruby method! Here’s the 3.3+ version - its pretty simple:
        def times
          #... a little C interop code
          i = 0
          while i < self
            yield i
            i = i.succ
          end
          self
        end

        Very simple. It’s just a basic while loop. Most importantly, it’s all Ruby code, which means YJIT should be able to introspect and optimize it!
        An aside on Integer#succ
        The slightly odd part of that code is i.succ. I’d never heard of Integer#succ, which apparently gives you the “successor” to an integer.


        I’ve never seen this show, and yet it’s the first thing I thought of when I learned about this method. Thanks, advertising.

        There was a PR to improve the performance of Integer#succ in early 2024, which helped me understand why anyone would ever use it:

        We use Integer#succ when we rewrite loop methods in Ruby (e.g. Integer#times and Array#each) because opt_succ (i = i.succ) is faster to dispatch on the interpreter than putobject 1; opt_plus (i += 1).
        https://github.com/ruby/ruby/pull/9519

        Integer#succ is like a virtual machine cheat code. It takes a common operation (adding 1 to an integer) and turns it from two virtual machine operations into one. We can call disasm on the times method to see that in action:
        puts RubyVM::InstructionSequence.disasm(1.method(:times))

        The Integer#times method gets broken down into a lot of Ruby VM bytecode, but we only care about a few lines:
        ...
        0025 getlocal_WC_0   i@0
        0027 opt_succ        <calldata!mid:succ, ARGS_SIMPLE>[CcCr]
        0029 setlocal_WC_0   i@0
        ...


        getlocal_WC_0 gets our i variable from the current scope. That’s the i in i.succ
        opt_succ performs the succ call in our i.succ. It will either call the actual Integer#succ method, or an optimized C function for small numbers
        In Ruby 3.4 with YJIT enabled, small numbers get optimized even further into machine code (just a note, not shown in the VM machine code)
        setlocal_WC_0 sets the result of opt_succ to our local variable i

        If we change from i = i.succ to i += 1, we now have two VM operations take the place of opt_succ:
        ...
        0025 getlocal_WC_0        i@0
        0027 putobject_INT2FIX_1_
        0028 opt_plus             <calldata!mid:+, argc:1, ARGS_SIMPLE>
        0029 setlocal_WC_0        i@0
        ...

        Everything is essentially the same as before, except now we have two steps to go through instead of one:

        putobject_INT2FIX_1_ pushes the integer 1 onto the virtual machine stack
        opt_plus is the + in our += 1, and calls either the Ruby + method or an optimized C function for small numbers
        There is probably a YJIT optimization for opt_plus as well

        If there is nothing else to learn from this code, it’s this: the kinds of optimizations you do at the VM and JIT level are deep. When writing general Ruby programs we typically don’t and shouldn’t consider the impact of one versus two machine code instructions. But at the JIT level, on the scale of millions and billions of operations, it matters!
        Back to Integer#times
        Let’s try running our benchmark code again, using times! Instead of iterating over ranges, we simply iterate for 10_000 and 100_000 times:
        u = ARGV[0].to_i        
        r = rand(10_000)        
        a = Array.new(10_000, 0)
        	
        10_000.times do |i|
          100_000.times do |j|
            a[i] += j % u
          end
          a[i] += r
        end
        	
        puts a[r]




        Loops




        Range#each
        25.57s


        Integer#times
        13.66s



        Nice! YJIT makes a much larger impact using Integer#times. That trims things down significantly, taking it down to 13.66 seconds on my machine. On @k0kobun’s machine it actually goes down to 9 seconds (and 8 seconds on Ruby 3.4).

        It’s probably Ruby 3.5’s job to make it faster than 8s though.

        We might look forward to even faster performance in Ruby 3.5. We’ll see!
        3. Array#each was converted from C to Ruby in Ruby 3.4
        CRuby continues to see C code rewritten in Ruby, and in Ruby 3.4 Array#each was one of those changes. Here is an example of the first attempt at implementing it:
        def each
          unless block_given?
            return to_enum(:each) { self.length }
          end
          i = 0
          while i < self.length
            yield self[i]
            i = i.succ
          end
          self
        end

        Super simple and readable! And YJIT optimizable!
        Unfortunately, due to something related to CRuby internals, it contained race conditions. A later implementation landed in Ruby 3.4.
        def each
          Primitive.attr! :inline_block, :c_trace

          unless defined?(yield)
            return Primitive.cexpr! 'SIZED_ENUMERATOR(self, 0, 0, ary_enum_length)'
          end
          _i = 0
          value = nil
          while Primitive.cexpr!(%q{ ary_fetch_next(self, LOCAL_PTR(_i), LOCAL_PTR(value)) })
            yield value
          end
          self
        end

        Unlike the first implementation, and unlike Integer#times, things are a bit more cryptic this time. This is definitely not pure Ruby code that anyone could be expected to write. Somehow, the Primitive module seems to allow evaluating C code from Ruby, and in doing so avoids the race conditions present in the pure Ruby solution3.
        By fetching indexes and values using C code, I think it results in a more atomic operation. I have no idea why the Primitive.cexpr! is used to return the enumerator, or what value Primitive.attr! :inline_block provides. Please comment if you have insights there!
        I was a little loose with my earlier Integer#times source code as well. That actually had a bit of this Primitive syntax as well. The core of the method is what we looked at, and it’s all Ruby, but the start of the method contains the same Primitive calls for :inline_block and returning the enumerator:
        def times
          Primitive.attr! :inline_block
          unless defined?(yield)
            return Primitive.cexpr! 'SIZED_ENUMERATOR(self, 0, 0, int_dotimes_size)'
          end
          #...

        Ok - it’s more cryptic than Integer#times was, but Array#each is mostly Ruby (on Ruby 3.4+). Let’s give it a try using arrays instead of ranges or times:
        u = ARGV[0].to_i
        r = rand(10_000)
        a = Array.new(10_000, 0)
        	
        outer = (0...10_000).to_a.freeze
        inner = (0...100_000).to_a.freeze
        outer.each do |i|
          inner.each do |j|
            a[i] += j % u
          end
          a[i] += r
        end
        	
        puts a[r]
        Despite the embedded C code, YJIT still seems capable of making some hefty performance optimizations. It’s within the same range as Integer#times!




        Loops




        Range#each
        25.57s


        Integer#times
        13.66s


        Array#each
        13.96s



        Microbenchmarking Ruby performance
        I’ve forked the original language implementation repo, and created my own repository called “Ruby Microbench”. It takes all of the examples discussed, as well as several other forms of doing the iteration in Ruby: https://github.com/jpcamara/ruby_microbench
        Here is the output of just running those using Ruby 3.4 with and without YJIT:




        fibonacciarray#each
        range#each
        times
        for
        while
        loop do




        Ruby 3.4 YJIT
        2.19s
        14.02s
        26.61s
        13.12s
        14.91s
        37.10s
        13.95s

        Ruby 3.4
        16.49s
        34.29s
        33.88s
        33.18s
        36.32s
        37.14s
        50.65s



        I have no idea why the while loop example I wrote seems to be so slow. I’d expected it to run much faster. Maybe there’s an issue with how I wrote it - feel free to open an issue or PR if you see something wrong with my implementation. The loop do (taken from @timtilberg’s example) runs around the same speed as Integer#times - although its performance is awful with YJIT turned off.

        📝 The for in loop and array#each have very similar performance, and that’s because at the Ruby VM bytecode level they are almost identical. for in is mostly syntactic sugar that transforms into an #each call in the VM. Thanks to Daniel Colson for pointing this out, and you can read his for loops in Ruby article for some additional information and nuance around for in!

        In addition to running Ruby 3.4, for fun I have it using rbenv to run:

        Ruby 3.3
        Ruby 3.3 YJIT
        Ruby 3.2
        Ruby 3.2 YJIT
        TruffleRuby 24.1
        Ruby Artichoke
        MRuby

        A few of the test runs are listed here:




        fibonacci

        array#each

        range#each
        times
        for
        while
        loop do




        Ruby 3.4 YJIT2.19s
        14.02s
        26.61s
        13.12s
        14.91s
        37.10s
        13.95s

        Ruby 3.4
        16.49s
        34.29s33.88s
        33.18s
        36.32s
        37.14s
        50.65s

        TruffleRuby 24.10.92s
        0.97s
        0.92s2.39s
        2.06s3.90s
        0.77s

        MRuby 3.3
        28.83s
        144.65s
        126.40s
        128.22s
        133.58s91.55s
        144.93s


        Artichoke
        19.71s
        236.10s214.55s
        214.51s
        215.95s
        174.70s
        264.67s



        Based on that, I’ve taken the original visualization and made a Ruby specific one here just for the fibonacci run:

        Speeding up range#each
        Can we, the non @k0kobun’s of the world, make range#each faster? If I monkey patch the Range class with a pure-ruby implementation, things do get much faster! Here’s my implementation:
        class Range
          def each
            beginning = self.begin
            ending = self.end
            i = beginning
            loop do
              break if i == ending
              yield i
              i = i.succ
            end
          end
        end

        And here is the change in performance - 2 seconds slower than times - not bad!



        				


        Time spent





        Range#each in C
        25.57s


        Range#each in Ruby
        16.64s



        This is obviously over-simplified. I don’t handle all of the different cases of Range, and there may be nuances I am missing. Also, most of the Ruby rewritten methods I’ve seen invoke a Primitive class for certain operations. I’d love to learn more about when and why it’s needed.
        But! It goes to show the power of moving things out of C and letting YJIT optimize our code. It can improve performance in ways that would be difficult or impossible to replicate in regular C code.
        YJIT standard library
        Last year Aaron Patterson wrote an article called Ruby Outperforms C, in which he rewrote a C extension in Ruby for some GraphQL parsing. The Ruby code outperformed C thanks to YJIT optimizations.
        This got me thinking that it would be interesting to see a kind of “YJIT standard library” emerge, where core ruby functionality run in C could be swapped out for Ruby implementations for use by people using YJIT.
        As it turns out, this is almost exactly what the core YJIT team has been doing. In many cases they’ve completely removed C code, but more recently they’ve created a with_yjit block. The code will only take effect if YJIT is enabled, and otherwise the C code will run. For example, this is howArray#each is implemented:
        with_yjit do
          if Primitive.rb_builtin_basic_definition_p(:each)
            undef :each

            def each # :nodoc:
              # ... we examined this code earlier ...
            end
          end
        end

        As of Ruby 3.3, YJIT can be lazily initialized. Thankfully the with_yjit code handles this - the appropriate with_yjit versions of methods will be run once YJIT is enabled:
        # Uses C-builtin
        [1, 2, 3].each do |i|
          puts i
        end

        RubyVM::YJIT.enable

        # Uses Ruby version, which can be YJIT optimized
        [1, 2, 3].each do |i|
          puts i
        end

        This is because with_yjit is a YJIT “hook”, which is called the moment YJIT is enabled. After being called, it is removed from the runtime using undef :with_yjit.
        Investigating YJIT optimizations
        We’ve looked at Ruby code. We’ve looked at C code. We’ve looked at Ruby VM bytecode. Why not take it one step deeper and look at some machine code? And maybe some Rust code? Hey - where are you going! Don’t walk away while I’m talking to you!
        If you haven’t walked away, or skipped to the next section, let’s take a look at a small sliver of YJIT while we’re here!
        We can see the machine code YJIT generates 😱. It’s possible by building CRuby from source with YJIT debug flags. If you’re on a Mac you can see my MacOS setup for hacking on CRuby or my docker setup for hacking on CRuby for more elaborate instructions on building Ruby. But the simplified step is when you go to ./configure Ruby, you hand in an option of --enable-yjit=dev:
        ./configure --enable-yjit=dev
        make install

        Let’s use our Integer#times example from earlier as our example Ruby code:
        u = ARGV[0].to_i
        r = rand(10_000)
        a = Array.new(10_000, 0)
        	
        10_000.times do |i|
          100_000.times do |j|
            a[i] += j % u
          end
          a[i] += r
        end
        	
        puts a[r]
        Because you’ve built Ruby with YJIT in dev mode, you can hand in the --yjit-dump-disasm flag when running your ruby program:
        ./ruby --yjit --yjit-dump-disasm test.rb 40

        Using this, we can see the machine code created. We’ll just focus in on one tiny part - the machine code equivalent of the Ruby VM bytecode we read earlier. Here is the original VM bytecode for opt_succ, which is generated when you call i.succ, the Integer#succ method:
        ...
        0027 opt_succ        <calldata!mid:succ, ARGS_SIMPLE>[CcCr]
        ...

        And here is the machine code YJIT generates in this scenario, on my Mac M2 arm64 architecture:
        # Block: times@<internal:numeric>:259 
        # reg_mapping: [Some(Stack(0)), None, None, None, None]
        # Insn: 0027 opt_succ (stack_size: 1)
        # call to Integer#succ
        # guard object is fixnum
        0x1096808c4: tst x1, #1
        0x1096808c8: b.eq #0x109683014
        0x1096808cc: nop 
        0x1096808d0: nop 
        0x1096808d4: nop 
        0x1096808d8: nop 
        0x1096808dc: nop 
        # Integer#succ
        0x1096808e0: adds x11, x1, #2
        0x1096808e4: b.vs #0x109683048
        0x1096808e8: mov x1, x11

        To be honest, I about 25% understand this, and 75% am combining my own logic and AI to learn it 🤫. Feel free to yell at me if I get it a little wrong, I’d love to learn more. But here’s how I break this down.
        # Block: times@<internal:numeric>:259

        👆🏼This roughly corresponds to the line i = i.succ in the Integer#times method in numeric.rb. I say roughly because in my current code I see that on line 258, but maybe it shows the end of the block it’s run in since YJIT compiles “blocks” of code:
        256: while i < self
        257:   yield i
        258:   i = i.succ
        259: end

        # reg_mapping: [Some(Stack(0)), None, None, None, None]
        # Insn: 0027 opt_succ (stack_size: 1)
        # call to Integer#succ

        👆🏼I have no idea what reg_mapping means - probably mapping how it uses a CPU register? Insn: 0027 opt_succ looks very familiar! That’s our VM bytecode! call to Integer#succ is just a helpful comment added. YJIT is capable of adding comments to the machine code. We still haven’t even left the safety of the comments 😅.
        # guard object is fixnum

        👆🏼This is interesting. I can find a corresponding bit of Rust code that maps directly to this. Let’s take a look at it:
        fn jit_rb_int_succ(
          //...
          asm: &mut Assembler,
          //...
        ) -> bool {
          // Guard the receiver is fixnum
          let recv_type = asm.ctx.get_opnd_type(StackOpnd(0));
          let recv = asm.stack_pop(1);
          if recv_type != Type::Fixnum {
            asm_comment!(asm, "guard object is fixnum");
            asm.test(recv, Opnd::Imm(RUBY_FIXNUM_FLAG as i64));
                 asm.jz(Target::side_exit(Counter::opt_succ_not_fixnum));
          }

          asm_comment!(asm, "Integer#succ");
          let out_val = asm.add(recv, Opnd::Imm(2)); // 2 is untagged Fixnum 1
          asm.jo(Target::side_exit(Counter::opt_succ_overflow));

          // Push the output onto the stack
          let dst = asm.stack_push(Type::Fixnum);
          asm.mov(dst, out_val);

          true
        }

        Oh nice! This is the actual YJIT Rust implementation of the opt_succ call. This is that optimization @k0kobun made to further improve opt_succ performance beyond the bytecode C function calls. We’re in the section that is checking if what we’re operating on is a Fixnum, which is a way small integers are stored internally in CRuby:
        if recv_type != TypeFixnum 
          asm_comment!(asm, "guard object is fixnum");
          asm.test(recv, Opnd::Imm(RUBY_FIXNUM_FLAG as i64));
          asm.jz(Target::side_exit(Counter::opt_succ_not_fixnum));
        }

        That becomes this machine code:
        # guard object is fixnum
        0x1096808c4: tst x1, #1
        0x1096808c8: b.eq #0x109683014

        asm.test generates tst x1, #1, which according to an AI bot I asked is checking the least significant bit, which is a Fixnum “tag” that indicates this is a Fixnum. If it’s Fixnum, the result is 1 and b.eq is false. If it’s not a Fixnum, the result is 0 and b.eq is true and jumps away from this code.
        0x1096808cc: nop 
        0x1096808d0: nop 
        0x1096808d4: nop 
        0x1096808d8: nop 
        0x1096808dc: nop 

        🤖 “NOPs for alignment/padding”. Thanks AI. I don’t know why it is needed, but at least I know what it probably is.
        Finally, we actually add 1 to the number.
        asm_comment!(asm, "Integer#succ");
        let out_val = asm.add(recv, Opnd::Imm(2)); // 2 is untagged Fixnum 1
        asm.jo(Target::side_exit(Counter::opt_succ_overflow));

        // Push the output onto the stack
        let dst = asm.stack_push(Type::Fixnum);
        asm.mov(dst, out_val);

        The Rust code generates our Integer#succ comment. Then, to add 1, because of the “Fixnum tag” data embedded within our integer, actually means we have to add 2 using adds x11, x1, #2 😵‍💫. If we overflow the space available, it exits to a different code path - b.vs is a branch on overflow. Otherwise, it stores the result with mov x1, x11!
        # Integer#succ
        0x1096808e0: adds x11, x1, #2
        0x1096808e4: b.vs #0x109683048
        0x1096808e8: mov x1, x11

        😮‍💨. That was a lot. And it seems like alot of working is being done, but because it’s such low level machine code it’s presumably super fast. We examined a teensy tiny portion of what YJIT is capable of generating - JITs are complicated!
        Thanks to @k0kobun for providing me with the commands and pointing me at the YJIT docs which contain tons of additional options as well.
        The future of CRuby optimizations
        The irony of language implementation is that you often work less in the language you’re implementing than you do in something lower-level - in Ruby’s case, that’s mostly C and some Rust.
        With a layer like YJIT, it potentially opens up a future where more of the language becomes plain Ruby, and Ruby developer contribution is easier. Many languages have a smaller low level core, and the majority of the language is written in itself (like Java, for instance). Maybe that’s a future for CRuby, someday! Until then, keep the YJIT optimizations coming, YJIT team!






        Naive in this case meaning that there are more efficient ways to calculate fibonacci numbers in a program ↩︎


        MJIT, the precursor to YJIT, made Ruby much faster on certain benchmarks. But on large realistic Rails applications it actually made things slower ↩︎


        When C code is running, it has to opt-in to releasing the GVL, so it’s more difficult for threads to corrupt or modify data mid-operation. The original Ruby version could yield the GVL at points that would invalidate the array. That’s my understanding of the situation anyways. ↩︎
---
