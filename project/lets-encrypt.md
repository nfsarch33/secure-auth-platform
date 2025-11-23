Based on comprehensive research for your Rakuten Symphony authentication project, here's a complete guide for implementing HTTPS/TLS in both local and development environments:

# Complete HTTPS/TLS Guide for Golang + React + PostgreSQL

## Overview: Two-Environment Strategy

Your authentication app needs HTTPS in two contexts:

1. **Local Development** (localhost) - Use **mkcert** for trusted self-signed certificates[letsencrypt**+1**](https://letsencrypt.org/docs/certificates-for-localhost/)
2. **Production/Staging** (public domain) - Use **Let's Encrypt** with automatic renewal[stackoverflow**+1**](https://stackoverflow.com/questions/37321760/how-to-set-up-lets-encrypt-for-a-go-server-application)

 **Critical: You CANNOT use Let's Encrypt for localhost** . Let's Encrypt requires publicly accessible domain validation.[dcycle**+1**](http://blog.dcycle.com/blog/2018-10-27/local-https-docker-compose/)

---

## Part 1: Local Development HTTPS (mkcert)

## Why mkcert?[tecmint**+1**](https://www.tecmint.com/mkcert-create-ssl-certs-for-local-development/)

* ✅  **No browser warnings** : Creates locally-trusted certificates
* ✅  **Zero configuration** : Automatic CA installation
* ✅  **Multi-domain support** : One cert for localhost, 127.0.0.1, custom domains
* ✅  **Cross-platform** : Works on macOS, Linux, Windows
* ✅  **Docker-friendly** : Certificates can be mounted as volumes

---

## Step 1: Install mkcert

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">bash</div></div><div><span><code><span><span class="token token"># macOS</span><span>
</span></span><span><span>brew </span><span class="token token">install</span><span> mkcert
</span></span><span><span>brew </span><span class="token token">install</span><span> nss  </span><span class="token token"># For Firefox support</span><span>
</span></span><span>
</span><span><span></span><span class="token token"># Linux (Ubuntu/Debian)</span><span>
</span></span><span><span></span><span class="token token">sudo</span><span></span><span class="token token">apt</span><span></span><span class="token token">install</span><span> libnss3-tools
</span></span><span><span></span><span class="token token">curl</span><span> -JLO </span><span class="token token">"https://dl.filippo.io/mkcert/latest?for=linux/amd64"</span><span>
</span></span><span><span></span><span class="token token">chmod</span><span> +x mkcert-v*-linux-amd64
</span></span><span><span></span><span class="token token">sudo</span><span></span><span class="token token">mv</span><span> mkcert-v*-linux-amd64 /usr/local/bin/mkcert
</span></span><span>
</span><span><span></span><span class="token token"># Verify installation</span><span>
</span></span><span>mkcert -version
</span><span></span></code></span></div></div></div></pre>

---

## Step 2: Create Local Certificate Authority

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">bash</div></div><div><span><code><span><span class="token token"># Create and trust local CA (one-time setup)</span><span>
</span></span><span>mkcert -install
</span><span>
</span><span><span></span><span class="token token"># Output:</span><span>
</span></span><span><span></span><span class="token token"># Created a new local CA at "/home/jason/.local/share/mkcert"</span><span>
</span></span><span><span></span><span class="token token"># The local CA is now installed in the system trust store!</span><span>
</span></span><span></span></code></span></div></div></div></pre>

This creates:

* `~/.local/share/mkcert/rootCA.pem` (root certificate)
* `~/.local/share/mkcert/rootCA-key.pem` (root private key)

 **⚠️ Security Note** : Never share `rootCA-key.pem` or use mkcert certs in production.[tecmint](https://www.tecmint.com/mkcert-create-ssl-certs-for-local-development/)

---

## Step 3: Generate Certificates for Your Project

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">bash</div></div><div><span><code><span><span class="token token"># Navigate to project root</span><span>
</span></span><span><span></span><span class="token token">cd</span><span> rakuten-fullstack-auth-challenge
</span></span><span>
</span><span><span></span><span class="token token"># Create certs directory</span><span>
</span></span><span><span></span><span class="token token">mkdir</span><span> -p certs
</span></span><span>
</span><span><span></span><span class="token token"># Generate certificates for multiple domains</span><span>
</span></span><span><span>mkcert </span><span class="token token punctuation">\</span><span>
</span></span><span><span>  -cert-file certs/local-cert.pem </span><span class="token token punctuation">\</span><span>
</span></span><span><span>  -key-file certs/local-key.pem </span><span class="token token punctuation">\</span><span>
</span></span><span><span>  localhost </span><span class="token token punctuation">\</span><span>
</span></span><span><span></span><span class="token token">127.0</span><span>.0.1 </span><span class="token token punctuation">\</span><span>
</span></span><span><span>  ::1 </span><span class="token token punctuation">\</span><span>
</span></span><span><span>  auth.local </span><span class="token token punctuation">\</span><span>
</span></span><span><span>  api.auth.local </span><span class="token token punctuation">\</span><span>
</span></span><span><span></span><span class="token token">"*.auth.local"</span><span>
</span></span><span>
</span><span><span></span><span class="token token"># Output:</span><span>
</span></span><span><span></span><span class="token token"># Created a new certificate valid for the following names 📜</span><span>
</span></span><span><span></span><span class="token token">#  - "localhost"</span><span>
</span></span><span><span></span><span class="token token">#  - "127.0.0.1"</span><span>
</span></span><span><span></span><span class="token token">#  - "::1"</span><span>
</span></span><span><span></span><span class="token token">#  - "auth.local"</span><span>
</span></span><span><span></span><span class="token token">#  - "api.auth.local"</span><span>
</span></span><span><span></span><span class="token token">#  - "*.auth.local"</span><span>
</span></span><span><span></span><span class="token token">#</span><span>
</span></span><span><span></span><span class="token token"># The certificate is at "certs/local-cert.pem" and the key at "certs/local-key.pem" ✅</span><span>
</span></span><span></span></code></span></div></div></div></pre>

---

## Step 4: Update /etc/hosts (Optional for Custom Domains)

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">bash</div></div><div><span><code><span><span class="token token"># Add custom local domains</span><span>
</span></span><span><span></span><span class="token token">sudo</span><span></span><span class="token token">nano</span><span> /etc/hosts
</span></span><span>
</span><span><span></span><span class="token token"># Add these lines:</span><span>
</span></span><span><span></span><span class="token token">127.0</span><span>.0.1   auth.local
</span></span><span><span></span><span class="token token">127.0</span><span>.0.1   api.auth.local
</span></span><span></span></code></span></div></div></div></pre>

---

## Step 5: Configure Golang Backend for HTTPS

 **backend/cmd/server/main.go** :

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">go</div></div><div><span><code><span><span class="token token">package</span><span> main
</span></span><span>
</span><span><span></span><span class="token token">import</span><span></span><span class="token token punctuation">(</span><span>
</span></span><span><span></span><span class="token token">"crypto/tls"</span><span>
</span></span><span><span></span><span class="token token">"log"</span><span>
</span></span><span><span></span><span class="token token">"net/http"</span><span>
</span></span><span><span></span><span class="token token">"os"</span><span>
</span></span><span>  
</span><span><span></span><span class="token token">"github.com/gin-gonic/gin"</span><span>
</span></span><span><span></span><span class="token token">"github.com/yourusername/rakuten-auth/internal/api"</span><span>
</span></span><span><span></span><span class="token token">"github.com/yourusername/rakuten-auth/internal/config"</span><span>
</span></span><span><span></span><span class="token token punctuation">)</span><span>
</span></span><span>
</span><span><span></span><span class="token token">func</span><span></span><span class="token token">main</span><span class="token token punctuation">(</span><span class="token token punctuation">)</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span>    cfg </span><span class="token token operator">:=</span><span> config</span><span class="token token punctuation">.</span><span class="token token">Load</span><span class="token token punctuation">(</span><span class="token token punctuation">)</span><span>
</span></span><span>  
</span><span><span>    router </span><span class="token token operator">:=</span><span> api</span><span class="token token punctuation">.</span><span class="token token">SetupRouter</span><span class="token token punctuation">(</span><span>cfg</span><span class="token token punctuation">)</span><span>
</span></span><span>  
</span><span><span></span><span class="token token">// Check if running in HTTPS mode</span><span>
</span></span><span><span></span><span class="token token">if</span><span> cfg</span><span class="token token punctuation">.</span><span>TLSEnabled </span><span class="token token punctuation">{</span><span>
</span></span><span><span>        log</span><span class="token token punctuation">.</span><span class="token token">Println</span><span class="token token punctuation">(</span><span class="token token">"Starting HTTPS server on :8443"</span><span class="token token punctuation">)</span><span>
</span></span><span>      
</span><span><span></span><span class="token token">// Configure TLS</span><span>
</span></span><span><span>        tlsConfig </span><span class="token token operator">:=</span><span></span><span class="token token operator">&</span><span>tls</span><span class="token token punctuation">.</span><span>Config</span><span class="token token punctuation">{</span><span>
</span></span><span><span>            MinVersion</span><span class="token token punctuation">:</span><span>               tls</span><span class="token token punctuation">.</span><span>VersionTLS12</span><span class="token token punctuation">,</span><span>
</span></span><span><span>            CurvePreferences</span><span class="token token punctuation">:</span><span></span><span class="token token punctuation">[</span><span class="token token punctuation">]</span><span>tls</span><span class="token token punctuation">.</span><span>CurveID</span><span class="token token punctuation">{</span><span>tls</span><span class="token token punctuation">.</span><span>CurveP521</span><span class="token token punctuation">,</span><span> tls</span><span class="token token punctuation">.</span><span>CurveP384</span><span class="token token punctuation">,</span><span> tls</span><span class="token token punctuation">.</span><span>CurveP256</span><span class="token token punctuation">}</span><span class="token token punctuation">,</span><span>
</span></span><span><span>            PreferServerCipherSuites</span><span class="token token punctuation">:</span><span></span><span class="token token boolean">true</span><span class="token token punctuation">,</span><span>
</span></span><span><span>            CipherSuites</span><span class="token token punctuation">:</span><span></span><span class="token token punctuation">[</span><span class="token token punctuation">]</span><span class="token token">uint16</span><span class="token token punctuation">{</span><span>
</span></span><span><span>                tls</span><span class="token token punctuation">.</span><span>TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384</span><span class="token token punctuation">,</span><span>
</span></span><span><span>                tls</span><span class="token token punctuation">.</span><span>TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256</span><span class="token token punctuation">,</span><span>
</span></span><span><span>                tls</span><span class="token token punctuation">.</span><span>TLS_RSA_WITH_AES_256_GCM_SHA384</span><span class="token token punctuation">,</span><span>
</span></span><span><span>                tls</span><span class="token token punctuation">.</span><span>TLS_RSA_WITH_AES_128_GCM_SHA256</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span>      
</span><span><span>        server </span><span class="token token operator">:=</span><span></span><span class="token token operator">&</span><span>http</span><span class="token token punctuation">.</span><span>Server</span><span class="token token punctuation">{</span><span>
</span></span><span><span>            Addr</span><span class="token token punctuation">:</span><span></span><span class="token token">":8443"</span><span class="token token punctuation">,</span><span>
</span></span><span><span>            Handler</span><span class="token token punctuation">:</span><span>   router</span><span class="token token punctuation">,</span><span>
</span></span><span><span>            TLSConfig</span><span class="token token punctuation">:</span><span> tlsConfig</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span>      
</span><span><span></span><span class="token token">// Start HTTPS server</span><span>
</span></span><span><span>        log</span><span class="token token punctuation">.</span><span class="token token">Fatal</span><span class="token token punctuation">(</span><span>server</span><span class="token token punctuation">.</span><span class="token token">ListenAndServeTLS</span><span class="token token punctuation">(</span><span>cfg</span><span class="token token punctuation">.</span><span>TLSCertFile</span><span class="token token punctuation">,</span><span> cfg</span><span class="token token punctuation">.</span><span>TLSKeyFile</span><span class="token token punctuation">)</span><span class="token token punctuation">)</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span></span><span class="token token">else</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span></span><span class="token token">// HTTP mode for testing</span><span>
</span></span><span><span>        log</span><span class="token token punctuation">.</span><span class="token token">Println</span><span class="token token punctuation">(</span><span class="token token">"Starting HTTP server on :8080"</span><span class="token token punctuation">)</span><span>
</span></span><span><span>        log</span><span class="token token punctuation">.</span><span class="token token">Fatal</span><span class="token token punctuation">(</span><span>router</span><span class="token token punctuation">.</span><span class="token token">Run</span><span class="token token punctuation">(</span><span class="token token">":8080"</span><span class="token token punctuation">)</span><span class="token token punctuation">)</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span></span></code></span></div></div></div></pre>

 **backend/internal/config/config.go** :

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">go</div></div><div><span><code><span><span class="token token">package</span><span> config
</span></span><span>
</span><span><span></span><span class="token token">import</span><span></span><span class="token token punctuation">(</span><span>
</span></span><span><span></span><span class="token token">"os"</span><span>
</span></span><span><span></span><span class="token token punctuation">)</span><span>
</span></span><span>
</span><span><span></span><span class="token token">type</span><span> Config </span><span class="token token">struct</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span>    DatabaseURL </span><span class="token token">string</span><span>
</span></span><span><span>    JWTSecret   </span><span class="token token">string</span><span>
</span></span><span><span>    TLSEnabled  </span><span class="token token">bool</span><span>
</span></span><span><span>    TLSCertFile </span><span class="token token">string</span><span>
</span></span><span><span>    TLSKeyFile  </span><span class="token token">string</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span>
</span><span><span></span><span class="token token">func</span><span></span><span class="token token">Load</span><span class="token token punctuation">(</span><span class="token token punctuation">)</span><span></span><span class="token token operator">*</span><span>Config </span><span class="token token punctuation">{</span><span>
</span></span><span><span></span><span class="token token">return</span><span></span><span class="token token operator">&</span><span>Config</span><span class="token token punctuation">{</span><span>
</span></span><span><span>        DatabaseURL</span><span class="token token punctuation">:</span><span></span><span class="token token">getEnv</span><span class="token token punctuation">(</span><span class="token token">"DATABASE_URL"</span><span class="token token punctuation">,</span><span></span><span class="token token">"postgres://auth:auth@localhost:5432/authdb?sslmode=disable"</span><span class="token token punctuation">)</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        JWTSecret</span><span class="token token punctuation">:</span><span></span><span class="token token">getEnv</span><span class="token token punctuation">(</span><span class="token token">"JWT_SECRET"</span><span class="token token punctuation">,</span><span></span><span class="token token">"your-secret-key-change-in-production"</span><span class="token token punctuation">)</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        TLSEnabled</span><span class="token token punctuation">:</span><span></span><span class="token token">getEnv</span><span class="token token punctuation">(</span><span class="token token">"TLS_ENABLED"</span><span class="token token punctuation">,</span><span></span><span class="token token">"true"</span><span class="token token punctuation">)</span><span></span><span class="token token operator">==</span><span></span><span class="token token">"true"</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        TLSCertFile</span><span class="token token punctuation">:</span><span></span><span class="token token">getEnv</span><span class="token token punctuation">(</span><span class="token token">"TLS_CERT_FILE"</span><span class="token token punctuation">,</span><span></span><span class="token token">"./certs/local-cert.pem"</span><span class="token token punctuation">)</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        TLSKeyFile</span><span class="token token punctuation">:</span><span></span><span class="token token">getEnv</span><span class="token token punctuation">(</span><span class="token token">"TLS_KEY_FILE"</span><span class="token token punctuation">,</span><span></span><span class="token token">"./certs/local-key.pem"</span><span class="token token punctuation">)</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span>
</span><span><span></span><span class="token token">func</span><span></span><span class="token token">getEnv</span><span class="token token punctuation">(</span><span>key</span><span class="token token punctuation">,</span><span> defaultValue </span><span class="token token">string</span><span class="token token punctuation">)</span><span></span><span class="token token">string</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span></span><span class="token token">if</span><span> value </span><span class="token token operator">:=</span><span> os</span><span class="token token punctuation">.</span><span class="token token">Getenv</span><span class="token token punctuation">(</span><span>key</span><span class="token token punctuation">)</span><span class="token token punctuation">;</span><span> value </span><span class="token token operator">!=</span><span></span><span class="token token">""</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span></span><span class="token token">return</span><span> value
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span><span></span><span class="token token">return</span><span> defaultValue
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span></span></code></span></div></div></div></pre>

---

## Step 6: Configure React Frontend for HTTPS

 **frontend/vite.config.ts** :[vitejs**+1**](https://v2.vitejs.dev/config/)

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">typescript</div></div><div><span><code><span><span class="token token">import</span><span></span><span class="token token punctuation">{</span><span> defineConfig </span><span class="token token punctuation">}</span><span></span><span class="token token">from</span><span></span><span class="token token">'vite'</span><span class="token token punctuation">;</span><span>
</span></span><span><span></span><span class="token token">import</span><span> react </span><span class="token token">from</span><span></span><span class="token token">'@vitejs/plugin-react'</span><span class="token token punctuation">;</span><span>
</span></span><span><span></span><span class="token token">import</span><span> fs </span><span class="token token">from</span><span></span><span class="token token">'fs'</span><span class="token token punctuation">;</span><span>
</span></span><span><span></span><span class="token token">import</span><span> path </span><span class="token token">from</span><span></span><span class="token token">'path'</span><span class="token token punctuation">;</span><span>
</span></span><span>
</span><span><span></span><span class="token token">export</span><span></span><span class="token token">default</span><span></span><span class="token token">defineConfig</span><span class="token token punctuation">(</span><span class="token token punctuation">{</span><span>
</span></span><span><span>  plugins</span><span class="token token operator">:</span><span></span><span class="token token punctuation">[</span><span class="token token">react</span><span class="token token punctuation">(</span><span class="token token punctuation">)</span><span class="token token punctuation">]</span><span class="token token punctuation">,</span><span>
</span></span><span>  
</span><span><span>  server</span><span class="token token operator">:</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span></span><span class="token token">// Enable HTTPS with mkcert certificates</span><span>
</span></span><span><span>    https</span><span class="token token operator">:</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span>      key</span><span class="token token operator">:</span><span> fs</span><span class="token token punctuation">.</span><span class="token token">readFileSync</span><span class="token token punctuation">(</span><span>path</span><span class="token token punctuation">.</span><span class="token token">resolve</span><span class="token token punctuation">(</span><span>__dirname</span><span class="token token punctuation">,</span><span></span><span class="token token">'../certs/local-key.pem'</span><span class="token token punctuation">)</span><span class="token token punctuation">)</span><span class="token token punctuation">,</span><span>
</span></span><span><span>      cert</span><span class="token token operator">:</span><span> fs</span><span class="token token punctuation">.</span><span class="token token">readFileSync</span><span class="token token punctuation">(</span><span>path</span><span class="token token punctuation">.</span><span class="token token">resolve</span><span class="token token punctuation">(</span><span>__dirname</span><span class="token token punctuation">,</span><span></span><span class="token token">'../certs/local-cert.pem'</span><span class="token token punctuation">)</span><span class="token token punctuation">)</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">,</span><span>
</span></span><span>  
</span><span><span></span><span class="token token">// Server configuration</span><span>
</span></span><span><span>    host</span><span class="token token operator">:</span><span></span><span class="token token">'localhost'</span><span class="token token punctuation">,</span><span>
</span></span><span><span>    port</span><span class="token token operator">:</span><span></span><span class="token token">3000</span><span class="token token punctuation">,</span><span>
</span></span><span>  
</span><span><span></span><span class="token token">// Proxy API requests to backend</span><span>
</span></span><span><span>    proxy</span><span class="token token operator">:</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span></span><span class="token token string-property property">'/api'</span><span class="token token operator">:</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span>        target</span><span class="token token operator">:</span><span></span><span class="token token">'https://localhost:8443'</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        changeOrigin</span><span class="token token operator">:</span><span></span><span class="token token boolean">true</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        secure</span><span class="token token operator">:</span><span></span><span class="token token boolean">false</span><span class="token token punctuation">,</span><span></span><span class="token token">// Accept self-signed certs in development</span><span>
</span></span><span><span></span><span class="token token function-variable">rewrite</span><span class="token token operator">:</span><span></span><span class="token token punctuation">(</span><span>path</span><span class="token token punctuation">)</span><span></span><span class="token token operator">=></span><span> path</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">,</span><span>
</span></span><span>  
</span><span><span></span><span class="token token">// Build configuration</span><span>
</span></span><span><span>  build</span><span class="token token operator">:</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span>    outDir</span><span class="token token operator">:</span><span></span><span class="token token">'dist'</span><span class="token token punctuation">,</span><span>
</span></span><span><span>    sourcemap</span><span class="token token operator">:</span><span></span><span class="token token boolean">true</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">)</span><span class="token token punctuation">;</span><span>
</span></span><span></span></code></span></div></div></div></pre>

 **Alternative: Using @vitejs/plugin-basic-ssl** :

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">bash</div></div><div><span><code><span><span class="token token">npm</span><span></span><span class="token token">install</span><span> --save-dev @vitejs/plugin-basic-ssl
</span></span><span></span></code></span></div></div></div></pre>

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">typescript</div></div><div><span><code><span><span class="token token">import</span><span></span><span class="token token punctuation">{</span><span> defineConfig </span><span class="token token punctuation">}</span><span></span><span class="token token">from</span><span></span><span class="token token">'vite'</span><span class="token token punctuation">;</span><span>
</span></span><span><span></span><span class="token token">import</span><span> react </span><span class="token token">from</span><span></span><span class="token token">'@vitejs/plugin-react'</span><span class="token token punctuation">;</span><span>
</span></span><span><span></span><span class="token token">import</span><span> basicSsl </span><span class="token token">from</span><span></span><span class="token token">'@vitejs/plugin-basic-ssl'</span><span class="token token punctuation">;</span><span>
</span></span><span>
</span><span><span></span><span class="token token">export</span><span></span><span class="token token">default</span><span></span><span class="token token">defineConfig</span><span class="token token punctuation">(</span><span class="token token punctuation">{</span><span>
</span></span><span><span>  plugins</span><span class="token token operator">:</span><span></span><span class="token token punctuation">[</span><span>
</span></span><span><span></span><span class="token token">react</span><span class="token token punctuation">(</span><span class="token token punctuation">)</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token">basicSsl</span><span class="token token punctuation">(</span><span class="token token punctuation">)</span><span class="token token punctuation">,</span><span></span><span class="token token">// Auto-generates self-signed cert</span><span>
</span></span><span><span></span><span class="token token punctuation">]</span><span class="token token punctuation">,</span><span>
</span></span><span>  
</span><span><span>  server</span><span class="token token operator">:</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span>    proxy</span><span class="token token operator">:</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span></span><span class="token token string-property property">'/api'</span><span class="token token operator">:</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span>        target</span><span class="token token operator">:</span><span></span><span class="token token">'https://localhost:8443'</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        changeOrigin</span><span class="token token operator">:</span><span></span><span class="token token boolean">true</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        secure</span><span class="token token operator">:</span><span></span><span class="token token boolean">false</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">)</span><span class="token token punctuation">;</span><span>
</span></span><span></span></code></span></div></div></div></pre>

---

## Step 7: Configure PostgreSQL with TLS

 **PostgreSQL SSL Configuration** :[sliplane**+2**](https://sliplane.io/blog/setup-tls-for-postgresql-in-docker)

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">bash</div></div><div><span><code><span><span class="token token"># Generate PostgreSQL server certificates</span><span>
</span></span><span><span></span><span class="token token">cd</span><span> certs
</span></span><span>
</span><span><span></span><span class="token token"># Generate server certificate (CN must match PostgreSQL hostname)</span><span>
</span></span><span><span>mkcert -cert-file postgres-cert.pem -key-file postgres-key.pem localhost </span><span class="token token">127.0</span><span>.0.1 postgres
</span></span><span>
</span><span><span></span><span class="token token"># Set correct permissions (PostgreSQL is strict about key permissions)</span><span>
</span></span><span><span></span><span class="token token">chmod</span><span></span><span class="token token">600</span><span> postgres-key.pem
</span></span><span><span></span><span class="token token">chmod</span><span></span><span class="token token">644</span><span> postgres-cert.pem
</span></span><span></span></code></span></div></div></div></pre>

 **postgres/ssl-config.sh** :

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">bash</div></div><div><span><code><span><span class="token token shebang important">#!/bin/bash</span><span>
</span></span><span><span></span><span class="token token">set</span><span> -e
</span></span><span>
</span><span><span></span><span class="token token"># Configure PostgreSQL to use SSL</span><span>
</span></span><span><span></span><span class="token token">echo</span><span></span><span class="token token">"Configuring PostgreSQL SSL..."</span><span>
</span></span><span>
</span><span><span></span><span class="token token">cat</span><span></span><span class="token token operator">>></span><span> /var/lib/postgresql/data/postgresql.conf </span><span class="token token operator"><<</span><span class="token token">EOF
</span></span><span class="token token">ssl = on
</span><span class="token token">ssl_cert_file = '/var/lib/postgresql/certs/postgres-cert.pem'
</span><span class="token token">ssl_key_file = '/var/lib/postgresql/certs/postgres-key.pem'
</span><span class="token token">ssl_prefer_server_ciphers = on
</span><span class="token token">ssl_min_protocol_version = 'TLSv1.2'
</span><span><span class="token token">EOF</span><span>
</span></span><span>
</span><span><span></span><span class="token token"># Update pg_hba.conf to require SSL</span><span>
</span></span><span><span></span><span class="token token">cat</span><span></span><span class="token token operator">>></span><span> /var/lib/postgresql/data/pg_hba.conf </span><span class="token token operator"><<</span><span class="token token">EOF
</span></span><span class="token token"># SSL connections
</span><span class="token token">hostssl all all 0.0.0.0/0 scram-sha-256
</span><span><span class="token token">EOF</span><span>
</span></span><span>
</span><span><span></span><span class="token token">echo</span><span></span><span class="token token">"PostgreSQL SSL configuration complete"</span><span>
</span></span><span></span></code></span></div></div></div></pre>

---

## Step 8: Docker Compose Configuration (Local Dev)

 **docker-compose.yml** :

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>version: '3.8'
</span></span><span>
</span><span>services:
</span><span>  postgres:
</span><span>    image: postgres:16-alpine
</span><span>    container_name: auth-postgres
</span><span>    environment:
</span><span>      POSTGRES_DB: authdb
</span><span>      POSTGRES_USER: auth
</span><span>      POSTGRES_PASSWORD: auth_password
</span><span>    ports:
</span><span>      - "5432:5432"
</span><span>    volumes:
</span><span>      - postgres_data:/var/lib/postgresql/data
</span><span>      - ./certs:/var/lib/postgresql/certs:ro
</span><span>      - ./postgres/ssl-config.sh:/docker-entrypoint-initdb.d/ssl-config.sh:ro
</span><span>    networks:
</span><span>      - auth-network
</span><span>    command: >
</span><span>      postgres
</span><span>      -c ssl=on
</span><span>      -c ssl_cert_file=/var/lib/postgresql/certs/postgres-cert.pem
</span><span>      -c ssl_key_file=/var/lib/postgresql/certs/postgres-key.pem
</span><span>
</span><span>  backend:
</span><span>    build:
</span><span>      context: ./backend
</span><span>      dockerfile: Dockerfile
</span><span>    container_name: auth-backend
</span><span>    environment:
</span><span>      DATABASE_URL: "postgres://auth:auth_password@postgres:5432/authdb?sslmode=require"
</span><span>      JWT_SECRET: "dev-secret-key"
</span><span>      TLS_ENABLED: "true"
</span><span>      TLS_CERT_FILE: "/app/certs/local-cert.pem"
</span><span>      TLS_KEY_FILE: "/app/certs/local-key.pem"
</span><span>    ports:
</span><span>      - "8443:8443"
</span><span>    volumes:
</span><span>      - ./certs:/app/certs:ro
</span><span>    depends_on:
</span><span>      - postgres
</span><span>    networks:
</span><span>      - auth-network
</span><span>
</span><span>  frontend:
</span><span>    build:
</span><span>      context: ./frontend
</span><span>      dockerfile: Dockerfile
</span><span>    container_name: auth-frontend
</span><span>    environment:
</span><span>      VITE_API_BASE_URL: "https://localhost:8443"
</span><span>    ports:
</span><span>      - "3000:3000"
</span><span>    volumes:
</span><span>      - ./certs:/app/certs:ro
</span><span>      - ./frontend/src:/app/src
</span><span>    depends_on:
</span><span>      - backend
</span><span>    networks:
</span><span>      - auth-network
</span><span>
</span><span>volumes:
</span><span>  postgres_data:
</span><span>
</span><span>networks:
</span><span>  auth-network:
</span><span>    driver: bridge
</span><span></span></code></span></div></div></div></pre>

---

## Step 9: Update Backend Dockerfile

 **backend/Dockerfile** :

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>FROM golang:1.21-alpine AS builder
</span></span><span>
</span><span>WORKDIR /app
</span><span>
</span><span># Copy go mod files
</span><span>COPY go.mod go.sum ./
</span><span>RUN go mod download
</span><span>
</span><span># Copy source code
</span><span>COPY . .
</span><span>
</span><span># Build
</span><span>RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server
</span><span>
</span><span># Runtime stage
</span><span>FROM alpine:latest
</span><span>
</span><span>RUN apk --no-cache add ca-certificates
</span><span>
</span><span>WORKDIR /root/
</span><span>
</span><span># Copy binary and certs directory structure
</span><span>COPY --from=builder /app/server .
</span><span>RUN mkdir -p /app/certs
</span><span>
</span><span>EXPOSE 8443
</span><span>
</span><span>CMD ["./server"]
</span><span></span></code></span></div></div></div></pre>

---

## Step 10: Testing Local HTTPS Setup

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">bash</div></div><div><span><code><span><span class="token token"># 1. Start Docker Compose</span><span>
</span></span><span><span></span><span class="token token">make</span><span> dev
</span></span><span><span></span><span class="token token"># or</span><span>
</span></span><span><span></span><span class="token token">docker-compose</span><span> up --build
</span></span><span>
</span><span><span></span><span class="token token"># 2. Access frontend</span><span>
</span></span><span><span></span><span class="token token"># Browser: https://localhost:3000 (no warnings!)</span><span>
</span></span><span>
</span><span><span></span><span class="token token"># 3. Test backend API</span><span>
</span></span><span><span></span><span class="token token">curl</span><span> https://localhost:8443/api/health
</span></span><span>
</span><span><span></span><span class="token token"># 4. Test PostgreSQL SSL connection</span><span>
</span></span><span><span>psql </span><span class="token token">"postgres://auth:auth_password@localhost:5432/authdb?sslmode=require"</span><span>
</span></span><span>
</span><span><span></span><span class="token token"># 5. Verify certificates</span><span>
</span></span><span>openssl s_client -connect localhost:8443 -showcerts
</span><span>
</span><span><span></span><span class="token token"># 6. Check PostgreSQL SSL</span><span>
</span></span><span><span>psql </span><span class="token token">"postgres://auth:auth_password@localhost:5432/authdb"</span><span></span><span class="token token punctuation">\</span><span>
</span></span><span><span>  -c </span><span class="token token">"SELECT * FROM pg_stat_ssl;"</span><span>
</span></span><span></span></code></span></div></div></div></pre>

---

## Part 2: Production HTTPS (Let's Encrypt with Traefik)

## Why Traefik?[github**+3**](https://github.com/bubelov/traefik-letsencrypt-compose)

* ✅  **Automatic Let's Encrypt** : Requests and renews certificates automatically
* ✅  **Docker-native** : Configuration via labels
* ✅  **Zero downtime** : Hot reload on configuration changes
* ✅  **HTTP → HTTPS redirect** : Automatic secure upgrade
* ✅  **Dashboard** : Built-in monitoring UI

---

## Step 1: Production Architecture

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>Internet
</span></span><span>    ↓
</span><span>Traefik (443/80)
</span><span>    ├─→ Frontend (React) → https://auth.yourdomain.com
</span><span>    ├─→ Backend (Golang) → https://api.auth.yourdomain.com
</span><span>    └─→ PostgreSQL (internal, TLS)
</span><span></span></code></span></div></div></div></pre>

---

## Step 2: Traefik Configuration

 **docker-compose.prod.yml** :

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>version: '3.8'
</span></span><span>
</span><span>services:
</span><span>  traefik:
</span><span>    image: traefik:v2.11
</span><span>    container_name: traefik
</span><span>    restart: unless-stopped
</span><span>    security_opt:
</span><span>      - no-new-privileges:true
</span><span>    command:
</span><span>      # API and dashboard
</span><span>      - "--api.dashboard=true"
</span><span>    
</span><span>      # Docker provider
</span><span>      - "--providers.docker=true"
</span><span>      - "--providers.docker.exposedbydefault=false"
</span><span>    
</span><span>      # Entrypoints
</span><span>      - "--entrypoints.web.address=:80"
</span><span>      - "--entrypoints.websecure.address=:443"
</span><span>    
</span><span>      # Let's Encrypt configuration
</span><span>      - "--certificatesresolvers.letsencrypt.acme.tlschallenge=true"
</span><span>      - "--certificatesresolvers.letsencrypt.acme.email=your-email@example.com"
</span><span>      - "--certificatesresolvers.letsencrypt.acme.storage=/letsencrypt/acme.json"
</span><span>    
</span><span>      # Uncomment for testing (staging environment)
</span><span>      # - "--certificatesresolvers.letsencrypt.acme.caserver=https://acme-staging-v02.api.letsencrypt.org/directory"
</span><span>    
</span><span>      # Logs
</span><span>      - "--log.level=INFO"
</span><span>      - "--accesslog=true"
</span><span>    ports:
</span><span>      - "80:80"
</span><span>      - "443:443"
</span><span>    volumes:
</span><span>      - /var/run/docker.sock:/var/run/docker.sock:ro
</span><span>      - traefik-certificates:/letsencrypt
</span><span>    networks:
</span><span>      - auth-network
</span><span>    labels:
</span><span>      # Enable Traefik for dashboard
</span><span>      - "traefik.enable=true"
</span><span>    
</span><span>      # Dashboard HTTP (redirect to HTTPS)
</span><span>      - "traefik.http.routers.dashboard-http.entrypoints=web"
</span><span>      - "traefik.http.routers.dashboard-http.rule=Host(`traefik.yourdomain.com`)"
</span><span>      - "traefik.http.routers.dashboard-http.middlewares=redirect-to-https"
</span><span>    
</span><span>      # Dashboard HTTPS
</span><span>      - "traefik.http.routers.dashboard.entrypoints=websecure"
</span><span>      - "traefik.http.routers.dashboard.rule=Host(`traefik.yourdomain.com`)"
</span><span>      - "traefik.http.routers.dashboard.service=api@internal"
</span><span>      - "traefik.http.routers.dashboard.tls.certresolver=letsencrypt"
</span><span>      - "traefik.http.routers.dashboard.middlewares=auth"
</span><span>    
</span><span>      # Middleware: HTTPS redirect
</span><span>      - "traefik.http.middlewares.redirect-to-https.redirectscheme.scheme=https"
</span><span>      - "traefik.http.middlewares.redirect-to-https.redirectscheme.permanent=true"
</span><span>    
</span><span>      # Middleware: Basic auth for dashboard
</span><span>      # Generate: htpasswd -nBC 10 admin
</span><span>      - "traefik.http.middlewares.auth.basicauth.users=admin:$$2y$$10$$zi5n43jq9S63gBqSJwHTH.nCai2vB0SW/ABPGg2jSGmJBVRo0A.ni"
</span><span>
</span><span>  postgres:
</span><span>    image: postgres:16-alpine
</span><span>    container_name: auth-postgres
</span><span>    restart: unless-stopped
</span><span>    environment:
</span><span>      POSTGRES_DB: authdb
</span><span>      POSTGRES_USER: auth
</span><span>      POSTGRES_PASSWORD_FILE: /run/secrets/db_password
</span><span>    secrets:
</span><span>      - db_password
</span><span>    volumes:
</span><span>      - postgres_data:/var/lib/postgresql/data
</span><span>      - ./postgres/ssl-config.sh:/docker-entrypoint-initdb.d/ssl-config.sh:ro
</span><span>      - ./certs/prod:/var/lib/postgresql/certs:ro
</span><span>    networks:
</span><span>      - auth-network
</span><span>    command: >
</span><span>      postgres
</span><span>      -c ssl=on
</span><span>      -c ssl_cert_file=/var/lib/postgresql/certs/server.crt
</span><span>      -c ssl_key_file=/var/lib/postgresql/certs/server.key
</span><span>
</span><span>  backend:
</span><span>    build:
</span><span>      context: ./backend
</span><span>      dockerfile: Dockerfile.prod
</span><span>    container_name: auth-backend
</span><span>    restart: unless-stopped
</span><span>    environment:
</span><span>      DATABASE_URL: "postgres://auth:${DB_PASSWORD}@postgres:5432/authdb?sslmode=require"
</span><span>      JWT_SECRET_FILE: /run/secrets/jwt_secret
</span><span>      TLS_ENABLED: "false"  # Traefik handles TLS termination
</span><span>    secrets:
</span><span>      - jwt_secret
</span><span>    depends_on:
</span><span>      - postgres
</span><span>    networks:
</span><span>      - auth-network
</span><span>    labels:
</span><span>      - "traefik.enable=true"
</span><span>    
</span><span>      # HTTP (redirect to HTTPS)
</span><span>      - "traefik.http.routers.backend-http.entrypoints=web"
</span><span>      - "traefik.http.routers.backend-http.rule=Host(`api.auth.yourdomain.com`)"
</span><span>      - "traefik.http.routers.backend-http.middlewares=redirect-to-https"
</span><span>    
</span><span>      # HTTPS
</span><span>      - "traefik.http.routers.backend.entrypoints=websecure"
</span><span>      - "traefik.http.routers.backend.rule=Host(`api.auth.yourdomain.com`)"
</span><span>      - "traefik.http.routers.backend.tls.certresolver=letsencrypt"
</span><span>      - "traefik.http.routers.backend.service=backend"
</span><span>      - "traefik.http.services.backend.loadbalancer.server.port=8080"
</span><span>
</span><span>  frontend:
</span><span>    build:
</span><span>      context: ./frontend
</span><span>      dockerfile: Dockerfile.prod
</span><span>      args:
</span><span>        VITE_API_BASE_URL: "https://api.auth.yourdomain.com"
</span><span>    container_name: auth-frontend
</span><span>    restart: unless-stopped
</span><span>    depends_on:
</span><span>      - backend
</span><span>    networks:
</span><span>      - auth-network
</span><span>    labels:
</span><span>      - "traefik.enable=true"
</span><span>    
</span><span>      # HTTP (redirect to HTTPS)
</span><span>      - "traefik.http.routers.frontend-http.entrypoints=web"
</span><span>      - "traefik.http.routers.frontend-http.rule=Host(`auth.yourdomain.com`)"
</span><span>      - "traefik.http.routers.frontend-http.middlewares=redirect-to-https"
</span><span>    
</span><span>      # HTTPS
</span><span>      - "traefik.http.routers.frontend.entrypoints=websecure"
</span><span>      - "traefik.http.routers.frontend.rule=Host(`auth.yourdomain.com`)"
</span><span>      - "traefik.http.routers.frontend.tls.certresolver=letsencrypt"
</span><span>      - "traefik.http.routers.frontend.service=frontend"
</span><span>      - "traefik.http.services.frontend.loadbalancer.server.port=80"
</span><span>
</span><span>volumes:
</span><span>  postgres_data:
</span><span>  traefik-certificates:
</span><span>
</span><span>networks:
</span><span>  auth-network:
</span><span>    driver: bridge
</span><span>
</span><span>secrets:
</span><span>  db_password:
</span><span>    file: ./secrets/db_password.txt
</span><span>  jwt_secret:
</span><span>    file: ./secrets/jwt_secret.txt
</span><span></span></code></span></div></div></div></pre>

---

## Step 3: Alternative - Golang with Built-in Let's Encrypt[github**+1**](https://github.com/gin-gonic/autotls)

 **If you prefer NOT using Traefik** , Golang can handle Let's Encrypt directly:

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">go</div></div><div><span><code><span><span class="token token">package</span><span> main
</span></span><span>
</span><span><span></span><span class="token token">import</span><span></span><span class="token token punctuation">(</span><span>
</span></span><span><span></span><span class="token token">"crypto/tls"</span><span>
</span></span><span><span></span><span class="token token">"log"</span><span>
</span></span><span><span></span><span class="token token">"net/http"</span><span>
</span></span><span>  
</span><span><span></span><span class="token token">"github.com/gin-gonic/gin"</span><span>
</span></span><span><span></span><span class="token token">"golang.org/x/crypto/acme/autocert"</span><span>
</span></span><span><span></span><span class="token token punctuation">)</span><span>
</span></span><span>
</span><span><span></span><span class="token token">func</span><span></span><span class="token token">main</span><span class="token token punctuation">(</span><span class="token token punctuation">)</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span>    router </span><span class="token token operator">:=</span><span> gin</span><span class="token token punctuation">.</span><span class="token token">Default</span><span class="token token punctuation">(</span><span class="token token punctuation">)</span><span>
</span></span><span>  
</span><span><span></span><span class="token token">// Your routes</span><span>
</span></span><span><span>    router</span><span class="token token punctuation">.</span><span class="token token">GET</span><span class="token token punctuation">(</span><span class="token token">"/api/health"</span><span class="token token punctuation">,</span><span></span><span class="token token">func</span><span class="token token punctuation">(</span><span>c </span><span class="token token operator">*</span><span>gin</span><span class="token token punctuation">.</span><span>Context</span><span class="token token punctuation">)</span><span></span><span class="token token punctuation">{</span><span>
</span></span><span><span>        c</span><span class="token token punctuation">.</span><span class="token token">JSON</span><span class="token token punctuation">(</span><span class="token token">200</span><span class="token token punctuation">,</span><span> gin</span><span class="token token punctuation">.</span><span>H</span><span class="token token punctuation">{</span><span class="token token">"status"</span><span class="token token punctuation">:</span><span></span><span class="token token">"healthy"</span><span class="token token punctuation">}</span><span class="token token punctuation">)</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">)</span><span>
</span></span><span>  
</span><span><span></span><span class="token token">// Let's Encrypt autocert manager</span><span>
</span></span><span><span>    certManager </span><span class="token token operator">:=</span><span> autocert</span><span class="token token punctuation">.</span><span>Manager</span><span class="token token punctuation">{</span><span>
</span></span><span><span>        Prompt</span><span class="token token punctuation">:</span><span>     autocert</span><span class="token token punctuation">.</span><span>AcceptTOS</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        HostPolicy</span><span class="token token punctuation">:</span><span> autocert</span><span class="token token punctuation">.</span><span class="token token">HostWhitelist</span><span class="token token punctuation">(</span><span class="token token">"api.auth.yourdomain.com"</span><span class="token token punctuation">)</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        Cache</span><span class="token token punctuation">:</span><span>      autocert</span><span class="token token punctuation">.</span><span class="token token">DirCache</span><span class="token token punctuation">(</span><span class="token token">"/var/www/.cache"</span><span class="token token punctuation">)</span><span class="token token punctuation">,</span><span></span><span class="token token">// Cache certs</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span>  
</span><span><span></span><span class="token token">// HTTPS server</span><span>
</span></span><span><span>    server </span><span class="token token operator">:=</span><span></span><span class="token token operator">&</span><span>http</span><span class="token token punctuation">.</span><span>Server</span><span class="token token punctuation">{</span><span>
</span></span><span><span>        Addr</span><span class="token token punctuation">:</span><span></span><span class="token token">":443"</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        Handler</span><span class="token token punctuation">:</span><span> router</span><span class="token token punctuation">,</span><span>
</span></span><span><span>        TLSConfig</span><span class="token token punctuation">:</span><span></span><span class="token token operator">&</span><span>tls</span><span class="token token punctuation">.</span><span>Config</span><span class="token token punctuation">{</span><span>
</span></span><span><span>            GetCertificate</span><span class="token token punctuation">:</span><span> certManager</span><span class="token token punctuation">.</span><span>GetCertificate</span><span class="token token punctuation">,</span><span>
</span></span><span><span>            MinVersion</span><span class="token token punctuation">:</span><span>     tls</span><span class="token token punctuation">.</span><span>VersionTLS12</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span class="token token punctuation">,</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span>  
</span><span><span></span><span class="token token">// HTTP server (for ACME challenges)</span><span>
</span></span><span><span></span><span class="token token">go</span><span> http</span><span class="token token punctuation">.</span><span class="token token">ListenAndServe</span><span class="token token punctuation">(</span><span class="token token">":80"</span><span class="token token punctuation">,</span><span> certManager</span><span class="token token punctuation">.</span><span class="token token">HTTPHandler</span><span class="token token punctuation">(</span><span class="token token boolean">nil</span><span class="token token punctuation">)</span><span class="token token punctuation">)</span><span>
</span></span><span>  
</span><span><span></span><span class="token token">// Start HTTPS server</span><span>
</span></span><span><span>    log</span><span class="token token punctuation">.</span><span class="token token">Fatal</span><span class="token token punctuation">(</span><span>server</span><span class="token token punctuation">.</span><span class="token token">ListenAndServeTLS</span><span class="token token punctuation">(</span><span class="token token">""</span><span class="token token punctuation">,</span><span></span><span class="token token">""</span><span class="token token punctuation">)</span><span class="token token punctuation">)</span><span>
</span></span><span><span></span><span class="token token punctuation">}</span><span>
</span></span><span></span></code></span></div></div></div></pre>

---

## Step 4: Deployment Script

 **scripts/deploy-production.sh** :

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">bash</div></div><div><span><code><span><span class="token token shebang important">#!/bin/bash</span><span>
</span></span><span><span></span><span class="token token">set</span><span> -e
</span></span><span>
</span><span><span></span><span class="token token">echo</span><span></span><span class="token token">"Deploying to production with Let's Encrypt..."</span><span>
</span></span><span>
</span><span><span></span><span class="token token"># Prerequisites check</span><span>
</span></span><span><span></span><span class="token token">if</span><span></span><span class="token token punctuation">[</span><span></span><span class="token token operator">!</span><span> -f .env.prod </span><span class="token token punctuation">]</span><span class="token token punctuation">;</span><span></span><span class="token token">then</span><span>
</span></span><span><span></span><span class="token token">echo</span><span></span><span class="token token">"Error: .env.prod not found"</span><span>
</span></span><span><span></span><span class="token token">exit</span><span></span><span class="token token">1</span><span>
</span></span><span><span></span><span class="token token">fi</span><span>
</span></span><span>
</span><span><span></span><span class="token token"># Load production environment variables</span><span>
</span></span><span><span></span><span class="token token">source</span><span> .env.prod
</span></span><span>
</span><span><span></span><span class="token token"># Create secrets directory</span><span>
</span></span><span><span></span><span class="token token">mkdir</span><span> -p secrets
</span></span><span><span></span><span class="token token">echo</span><span></span><span class="token token">"</span><span class="token token">$DB_PASSWORD</span><span class="token token">"</span><span></span><span class="token token operator">></span><span> secrets/db_password.txt
</span></span><span><span></span><span class="token token">echo</span><span></span><span class="token token">"</span><span class="token token">$JWT_SECRET</span><span class="token token">"</span><span></span><span class="token token operator">></span><span> secrets/jwt_secret.txt
</span></span><span>
</span><span><span></span><span class="token token"># Generate PostgreSQL certificates (production)</span><span>
</span></span><span><span></span><span class="token token">cd</span><span> certs/prod
</span></span><span><span>openssl req -new -x509 -days </span><span class="token token">365</span><span> -nodes -text </span><span class="token token punctuation">\</span><span>
</span></span><span><span>  -out server.crt </span><span class="token token punctuation">\</span><span>
</span></span><span><span>  -keyout server.key </span><span class="token token punctuation">\</span><span>
</span></span><span><span>  -subj </span><span class="token token">"/CN=postgres"</span><span>
</span></span><span><span></span><span class="token token">chmod</span><span></span><span class="token token">600</span><span> server.key
</span></span><span><span></span><span class="token token">chmod</span><span></span><span class="token token">644</span><span> server.crt
</span></span><span><span></span><span class="token token">cd</span><span></span><span class="token token punctuation">..</span><span>/</span><span class="token token punctuation">..</span><span>
</span></span><span>
</span><span><span></span><span class="token token"># Initialize Let's Encrypt certificates directory</span><span>
</span></span><span><span></span><span class="token token">mkdir</span><span> -p traefik-letsencrypt
</span></span><span><span></span><span class="token token">touch</span><span> traefik-letsencrypt/acme.json
</span></span><span><span></span><span class="token token">chmod</span><span></span><span class="token token">600</span><span> traefik-letsencrypt/acme.json
</span></span><span>
</span><span><span></span><span class="token token"># Deploy with Docker Compose</span><span>
</span></span><span><span></span><span class="token token">docker-compose</span><span> -f docker-compose.prod.yml up -d
</span></span><span>
</span><span><span></span><span class="token token">echo</span><span></span><span class="token token">"✅ Production deployment complete!"</span><span>
</span></span><span><span></span><span class="token token">echo</span><span></span><span class="token token">"Frontend: https://auth.yourdomain.com"</span><span>
</span></span><span><span></span><span class="token token">echo</span><span></span><span class="token token">"Backend: https://api.auth.yourdomain.com"</span><span>
</span></span><span><span></span><span class="token token">echo</span><span></span><span class="token token">"Traefik Dashboard: https://traefik.yourdomain.com"</span><span>
</span></span><span></span></code></span></div></div></div></pre>

---

## Part 3: Makefile Integration

 **Makefile** :

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">makefile</div></div><div><span><code><span><span class="token token builtin-target">.PHONY</span><span class="token token punctuation">:</span><span> help certs-local dev prod
</span></span><span>
</span><span><span></span><span class="token token target">help</span><span class="token token punctuation">:</span><span></span><span class="token token">## Show this help</span><span>
</span></span><span><span>	</span><span class="token token operator">@</span><span>grep -E </span><span class="token token">'^[a-zA-Z_-]+:.*?## .*$$'</span><span></span><span class="token token">$</span><span class="token token punctuation">(</span><span>MAKEFILE_LIST</span><span class="token token punctuation">)</span><span></span><span class="token token operator">|</span><span> sort </span><span class="token token operator">|</span><span> awk 'BEGIN </span><span class="token token punctuation">{</span><span>FS </span><span class="token token operator">=</span><span></span><span class="token token">":.*?## "</span><span class="token token punctuation">}</span><span class="token token punctuation">;</span><span></span><span class="token token punctuation">{</span><span>printf "\033
</span></span><span>
</span><span>---
</span><span>
</span><span><span></span><span class="token token">## Part 4: Environment Configuration</span><span>
</span></span><span>
</span><span><span></span><span class="token token target">**.env.local**</span><span class="token token punctuation">:</span><span>
</span></span><span>
</span><span></span></code></span></div></div></div></pre>

# Local development with mkcert

TLS_ENABLED=true

TLS_CERT_FILE=./certs/local-cert.pem

TLS_KEY_FILE=./certs/local-key.pem

DATABASE_URL=postgres://auth:auth_password@localhost:5432/authdb?sslmode=require

JWT_SECRET=dev-secret-key-change-in-production

VITE_API_BASE_URL=[https://localhost:8443](https://localhost:8443/)

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>
</span></span><span>**.env.prod**:
</span><span>
</span><span></span></code></span></div></div></div></pre>

# Production with Let's Encrypt (via Traefik)

TLS_ENABLED=false  # Traefik handles TLS termination

DATABASE_URL=postgres://auth:${DB_PASSWORD}@postgres:5432/authdb?sslmode=verify-full

DB_PASSWORD=your-secure-password-here

JWT_SECRET=your-secure-jwt-secret-here

DOMAIN=yourdomain.com

EMAIL=[your-email@yourdomain.com](mailto:your-email@yourdomain.com)

VITE_API_BASE_URL=[https://api.auth.yourdomain.com](https://api.auth.yourdomain.com/)

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>
</span></span><span>---
</span><span>
</span><span>## Part 5: Testing Checklist
</span><span>
</span><span>### Local Development Testing
</span><span>
</span><span></span></code></span></div></div></div></pre>

# 1. Generate certificates

make certs-local

# 2. Start services

make dev

# 3. Test backend HTTPS

curl [https://localhost:8443/api/health](https://localhost:8443/api/health)

# Expected:

# 4. Test frontend (browser)

# Open: [https://localhost:3000](https://localhost:3000/)

# Expected: No SSL warnings

# 5. Test PostgreSQL SSL

docker exec -it auth-postgres psql -U auth -d authdb -c "SELECT * FROM pg_stat_ssl;"

# Expected: ssl | t (true)

# 6. Verify certificate chain

openssl s_client -connect localhost:8443 -showcerts

# 7. Test API authentication flow

curl -X POST [https://localhost:8443/api/auth/signup](https://localhost:8443/api/auth/signup)

-H "Content-Type: application/json"

-d '{"email":"[test@example.com](mailto:test@example.com)","password":"SecurePass123!"}'

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>
</span></span><span>### Production Testing
</span><span>
</span><span></span></code></span></div></div></div></pre>

# 1. Deploy

make prod

# 2. Wait for Let's Encrypt (30-60 seconds)

# Monitor: docker logs traefik

# 3. Test HTTPS endpoints

curl [https://api.auth.yourdomain.com/api/health](https://api.auth.yourdomain.com/api/health)

# 4. Verify certificate

echo | openssl s_client -servername api.auth.yourdomain.com -connect api.auth.yourdomain.com:443 2>/dev/null | openssl x509 -noout -issuer -dates

# Expected issuer: Let's Encrypt Authority X3

# 5. Test SSL Labs score

# Visit: [https://www.ssllabs.com/ssltest/analyze.html?d=api.auth.yourdomain.com](https://www.ssllabs.com/ssltest/analyze.html?d=api.auth.yourdomain.com)

# Expected: A or A+ grade

# 6. Test automatic HTTP → HTTPS redirect

curl -I [http://auth.yourdomain.com](http://auth.yourdomain.com/)

# Expected: 301 Moved Permanently, Location: https://...

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>
</span></span><span>---
</span><span>
</span><span>## Part 6: Troubleshooting
</span><span>
</span><span>### Common Issues
</span><span>
</span><span>**1. "x509: certificate signed by unknown authority"**
</span><span>
</span><span></span></code></span></div></div></div></pre>

# Solution: Trust mkcert CA

mkcert -install

# Verify CA installation

ls ~/.local/share/mkcert/

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>
</span></span><span>**2. PostgreSQL SSL connection fails**
</span><span>
</span><span></span></code></span></div></div></div></pre>

# Check certificate permissions

ls -l certs/postgres-*.pem

# Expected: -rw------- (600) for key, -rw-r--r-- (644) for cert

# Fix permissions

chmod 600 certs/postgres-key.pem

chmod 644 certs/postgres-cert.pem

# Restart PostgreSQL

docker-compose restart postgres

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>
</span></span><span>**3. Let's Encrypt rate limits**
</span><span>
</span><span></span></code></span></div></div></div></pre>

# Use staging environment for testing

# In docker-compose.prod.yml, uncomment:

# - "--certificatesresolvers.letsencrypt.acme.caserver=[https://acme-staging-v02.api.letsencrypt.org/directory](https://acme-staging-v02.api.letsencrypt.org/directory)"

# After testing, remove line and redeploy

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>
</span></span><span>**4. Vite proxy not working with HTTPS backend**
</span><span>
</span><span></span></code></span></div></div></div></pre>

// vite.config.ts - Set secure: false

proxy: {

'/api': {

target: '[https://localhost:8443](https://localhost:8443/)',

changeOrigin: true,

secure: false, // Accept self-signed certs

},

},

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>
</span></span><span>**5. Browser still shows warning**
</span><span>
</span><span></span></code></span></div></div></div></pre>

# Chrome: Visit chrome://flags/#allow-insecure-localhost

# Enable "Allow invalid certificates for resources loaded from localhost"

# Or regenerate certificates including specific hostname

mkcert -cert-file certs/local-cert.pem -key-file certs/local-key.pem $(hostname) localhost 127.0.0.1

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>
</span></span><span>---
</span><span>
</span><span>## Summary: Quick Start Commands
</span><span>
</span><span></span></code></span></div></div></div></pre>

# Local development (mkcert)

make certs-local    # Generate local certificates (one-time)

make dev            # Start HTTPS development environment

# Production (Let's Encrypt via Traefik)

make prod           # Deploy with automatic Let's Encrypt

# Testing

make test-https     # Test all HTTPS endpoints

make ssl-info       # Show SSL certificate details

<pre class="not-prose w-full rounded font-mono text-sm font-extralight"><div class="codeWrapper text-light selection:text-super selection:bg-super/10 my-md relative flex flex-col rounded-lg font-mono text-sm font-normal bg-subtler"><div class="translate-y-xs -translate-x-xs bottom-xl mb-xl flex h-0 items-start justify-end md:sticky md:top-[calc(var(--header-height)+var(--size-xs))]"><div class="overflow-hidden rounded-full border-subtlest ring-subtlest divide-subtlest bg-base"><div class="border-subtlest ring-subtlest divide-subtlest bg-subtler"><button data-testid="copy-code-button" aria-label="Copy code" type="button" class="focus-visible:bg-subtle hover:bg-subtle text-quiet  hover:text-foreground dark:hover:bg-subtle font-sans focus:outline-none outline-none outline-transparent transition duration-300 ease-out select-none items-center relative group/button font-semimedium justify-center text-center items-center rounded-full cursor-pointer active:scale-[0.97] active:duration-150 active:ease-outExpo origin-center whitespace-nowrap inline-flex text-sm h-8 aspect-square" data-state="closed"><div class="flex items-center min-w-0 gap-two justify-center"><div class="flex shrink-0 items-center justify-center size-4"><svg role="img" class="inline-flex fill-current" width="16" height="16"><use xlink:href="#pplx-icon-copy"></use></svg></div></div></button></div></div></div><div class="-mt-xl"><div><div data-testid="code-language-indicator" class="text-quiet bg-subtle py-xs px-sm inline-block rounded-br rounded-tl-lg text-xs font-thin">text</div></div><div><span><code><span><span>
</span></span><span>---
</span><span>
</span><span>## Security Best Practices
</span><span>
</span><span>1. ✅ **Never commit certificates** - Add `certs/` to `.gitignore`
</span><span>2. ✅ **Use secrets management** - Store `JWT_SECRET`, `DB_PASSWORD` in Docker secrets
</span><span>3. ✅ **Minimum TLS 1.2** - Enforce in both Golang and PostgreSQL
</span><span>4. ✅ **Strong cipher suites** - Configure in `tls.Config`
</span><span>5. ✅ **HTTP → HTTPS redirect** - Enforce in production (Traefik handles automatically)
</span><span>6. ✅ **Certificate renewal** - Let's Encrypt auto-renews every 60 days (Traefik handles)
</span><span>7. ✅ **Monitor expiry** - Set up alerts for certificate expiration
</span><span>
</span><span>---
</span><span>
</span><span>This complete guide provides production-grade HTTPS for both local development (mkcert) and production deployment (Let's Encrypt with Traefik), tailored specifically for your Golang + React + PostgreSQL authentication application.
</span><span></span></code></span></div></div></div></pre>

1. [https://letsencrypt.org/docs/certificates-for-localhost/](https://letsencrypt.org/docs/certificates-for-localhost/)
2. [https://www.tecmint.com/mkcert-create-ssl-certs-for-local-development/](https://www.tecmint.com/mkcert-create-ssl-certs-for-local-development/)
3. [https://stackoverflow.com/questions/37321760/how-to-set-up-lets-encrypt-for-a-go-server-application](https://stackoverflow.com/questions/37321760/how-to-set-up-lets-encrypt-for-a-go-server-application)
4. [https://github.com/gin-gonic/autotls](https://github.com/gin-gonic/autotls)
5. [http://blog.dcycle.com/blog/2018-10-27/local-https-docker-compose/](http://blog.dcycle.com/blog/2018-10-27/local-https-docker-compose/)
6. [https://github.com/FiloSottile/mkcert/issues/349](https://github.com/FiloSottile/mkcert/issues/349)
7. [https://dev.to/aschmelyun/using-the-magic-of-mkcert-to-enable-valid-https-on-local-dev-sites-3a3c](https://dev.to/aschmelyun/using-the-magic-of-mkcert-to-enable-valid-https-on-local-dev-sites-3a3c)
8. [https://v2.vitejs.dev/config/](https://v2.vitejs.dev/config/)
9. [https://vite.dev/config/server-options](https://vite.dev/config/server-options)
10. [https://sliplane.io/blog/setup-tls-for-postgresql-in-docker](https://sliplane.io/blog/setup-tls-for-postgresql-in-docker)
11. [https://www.crunchydata.com/blog/ssl-certificate-authentication-postgresql-docker-containers](https://www.crunchydata.com/blog/ssl-certificate-authentication-postgresql-docker-containers)
12. [https://dev.to/whchi/setting-up-tls-connection-for-containerized-postgresql-database-1kmh](https://dev.to/whchi/setting-up-tls-connection-for-containerized-postgresql-database-1kmh)
13. [https://github.com/bubelov/traefik-letsencrypt-compose](https://github.com/bubelov/traefik-letsencrypt-compose)
14. [https://jnidzwetzki.github.io/2022/08/27/https-reverse-proxy-with-docker-traefik-and-lets-encrypt.html](https://jnidzwetzki.github.io/2022/08/27/https-reverse-proxy-with-docker-traefik-and-lets-encrypt.html)
15. [https://cylab.be/blog/258/automatic-docker-reverse-proxy-with-traefik](https://cylab.be/blog/258/automatic-docker-reverse-proxy-with-traefik)
16. [https://www.damirscorner.com/blog/posts/20250321-TraefikReverseProxyForDocker.html](https://www.damirscorner.com/blog/posts/20250321-TraefikReverseProxyForDocker.html)
17. [https://docs.vultr.com/secure-a-golang-web-server-with-a-selfsigned-or-lets-encrypt-ssl-certificate](https://docs.vultr.com/secure-a-golang-web-server-with-a-selfsigned-or-lets-encrypt-ssl-certificate)
18. [https://knplabs.com/en/blog/how-to-handle-https-with-docker-compose-and-mkcert-for-local-development](https://knplabs.com/en/blog/how-to-handle-https-with-docker-compose-and-mkcert-for-local-development)
19. [https://www.dsebastien.net/2020-01-08-deploying-tls-certificates-for-local-development-and-production-using-kubernetes-cert-manager/](https://www.dsebastien.net/2020-01-08-deploying-tls-certificates-for-local-development-and-production-using-kubernetes-cert-manager/)
20. [https://calmops.com/golang/how-to-create-a-https-and-tls-server-in-go/](https://calmops.com/golang/how-to-create-a-https-and-tls-server-in-go/)
21. [https://dev.to/vishalraj82/using-https-in-docker-for-local-development-nc7](https://dev.to/vishalraj82/using-https-in-docker-for-local-development-nc7)
22. [https://learn.microsoft.com/en-us/aspnet/core/security/docker-compose-https?view=aspnetcore-9.0](https://learn.microsoft.com/en-us/aspnet/core/security/docker-compose-https?view=aspnetcore-9.0)
23. [https://www.reddit.com/r/homelab/comments/z43334/how_to_create_ssl_certs_for_local_domain/](https://www.reddit.com/r/homelab/comments/z43334/how_to_create_ssl_certs_for_local_domain/)
24. [https://community.letsencrypt.org/t/how-to-enable-lets-encrypt-tls-certification-on-my-own-golang-project/184353](https://community.letsencrypt.org/t/how-to-enable-lets-encrypt-tls-certification-on-my-own-golang-project/184353)
25. [https://smithtalkstech.com/2018/04/25/creating-a-self-signed-ssl-certificate-for-local-docker-development/](https://smithtalkstech.com/2018/04/25/creating-a-self-signed-ssl-certificate-for-local-docker-development/)
26. [https://letsencrypt.org](https://letsencrypt.org/)
27. [https://stackoverflow.com/questions/57142452/adding-ssl-certificate-for-local-development-with-docker](https://stackoverflow.com/questions/57142452/adding-ssl-certificate-for-local-development-with-docker)
28. [https://community.letsencrypt.org/t/certificate-for-development-environment/121224](https://community.letsencrypt.org/t/certificate-for-development-environment/121224)
29. [https://www.reddit.com/r/golang/comments/91i9gn/go_server_best_practices_for_https/](https://www.reddit.com/r/golang/comments/91i9gn/go_server_best_practices_for_https/)
30. [https://dev.to/ghacosta/til-setting-up-proxy-server-on-vite-2cng](https://dev.to/ghacosta/til-setting-up-proxy-server-on-vite-2cng)
31. [https://www.geeksforgeeks.org/reactjs/how-to-configure-proxy-in-vite/](https://www.geeksforgeeks.org/reactjs/how-to-configure-proxy-in-vite/)
32. [https://gist.github.com/mrw34/c97bb03ea1054afb551886ffc8b63c3b](https://gist.github.com/mrw34/c97bb03ea1054afb551886ffc8b63c3b)
33. [https://github.com/vitejs/vite/issues/564](https://github.com/vitejs/vite/issues/564)
34. [https://stackoverflow.com/questions/55072221/deploying-postgresql-docker-with-ssl-certificate-and-key-with-volumes](https://stackoverflow.com/questions/55072221/deploying-postgresql-docker-with-ssl-certificate-and-key-with-volumes)
35. [https://doc.traefik.io/traefik/expose/docker/](https://doc.traefik.io/traefik/expose/docker/)
36. [https://stackoverflow.com/questions/64677212/how-to-configure-proxy-in-vite](https://stackoverflow.com/questions/64677212/how-to-configure-proxy-in-vite)
37. [https://www.red-gate.com/simple-talk/?p=107543](https://www.red-gate.com/simple-talk/?p=107543)
38. [https://www.tva.sg/traefik-reverse-proxy-the-complete-self-hosting-guide-for-https-and-ssl-automation/](https://www.tva.sg/traefik-reverse-proxy-the-complete-self-hosting-guide-for-https-and-ssl-automation/)
39. [https://www.reddit.com/r/reactjs/comments/wxi26o/having_problem_configuring_proxy_with_vite/](https://www.reddit.com/r/reactjs/comments/wxi26o/having_problem_configuring_proxy_with_vite/)
40. [https://maplibre.org/martin/pg-ssl-certificates.html](https://maplibre.org/martin/pg-ssl-certificates.html)
