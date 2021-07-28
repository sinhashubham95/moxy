# Moxy

[![GoDoc](https://godoc.org/github.com/sinhashubham95/moxy?status.svg)](https://pkg.go.dev/github.com/sinhashubham95/moxy)
[![Release](https://img.shields.io/github/v/release/sinhashubham95/moxy?sort=semver)](https://github.com/sinhashubham95/moxy/releases)
[![Report](https://goreportcard.com/badge/github.com/sinhashubham95/moxy)](https://goreportcard.com/report/github.com/sinhashubham95/moxy)
[![Coverage Status](https://coveralls.io/repos/github/sinhashubham95/moxy/badge.svg?branch=master)](https://coveralls.io/github/sinhashubham95/moxy?branch=master)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#server-applications)

Moxy is a simple mocker and proxy application server. Here you can create mock endpoints as well as proxy requests in case no mock exists for the endpoint.

## How it works

![Moxy Architecture Diagram](./moxy.png)

## Features

### Small, Pragmatic and Easy to Use

- Dockerized

- Compiled

- Easily configurable via Environment Variables

- Self-contained, does not require Go or any other dependency, just run the binary or the container

### File-based Persistence

- No heavy database involved.

- Saves the mock endpoints in files.

### Security

- TLS can be enabled by providing your own SSL/TLS Certificates.

### Reliability

- Uses [Go Actuator](https://github.com/sinhashubham95/go-actuator)

- Fully Tested, Unit, Functional & Linted & 0 Race Conditions Detected.

### Customizations

- Mock endpoints created are highly customizable.

- Application port can be configured via the environment variable.

- Database file path can be configured via the environment variable.

## Project Versioning

Moxy uses [semantic versioning](http://semver.org/). No API changes will be introduced in the minor and patch version changes. New minor versions might add additional features to the API.

## Getting Started(standalone application)

You can download a pre-compiled package for Linux or MAC OSX [here](https://github.com/sinhashubham95/moxy/releases/latest).

You can also pull the latest docker image for Moxy from [Docker Hub](https://github.com/sinhashubham95/moxy/pkgs/container/moxy).

```shell
docker pull ghcr.io/sinhashubham95/moxy:latest
```

Create an empty directory, change into it and run the following to start Moxy:

```shell
docker run --rm --user $(id -u):$(id -g) -v $PWD:/data -p 9091:9091 sinhashubham95/moxy -host 0.0.0.0
```

The container will have access to the current local directory and all sub-folders.

## Usage

### Mock - POST /moxy/mock

This is used to create a new mock endpoint. Each mock endpoint created can be tagged, which means you can create multiple mocks for the same endpoint.

```http request
POST /moxy/mock

{
    "tag": "",
    "method": "",
    "path": "",
    "responseDelayInMillis": 0,
    "responseStatus": 0,
    "responseBody": {}
}
```

1. **Tag**: This is used to tag the created mock endpoint. It is used to set the context for this endpoint. You can have many instances of the same endpoint.

2. **Method**: This is the request method for this endpoint. It can either be `GET`, `POST`, `PUT ` or `DELETE`.

3. **Path**: This is the request path for this endpoint. Note that the request path should not start with either `/actuator` or `/moxy`.

4. **Response Delay In Millis**: This is the delay added before actually returned the mocked response. Note that this delay is over and above the delay added over the network. In other words, this is the server side delay. It can be useful in testing the resiliency of the applications using this endpoint and testing the timeout functionality.

5. **Response Status**: This is the status code for the mocked response. It should be in the range of a valid HTTP status code which is `>=100` and `<= 599`.

6. **Response Body**: This is the response body for the mocked response. If it is a JSON, it can be provided as a JSON body itself, but anything apart from JSON has to be converted to a string and sent. Don't worry the mocked response will be returned with the correct type, that will be taken care of automatically. 

### Remove Mock - DELETE /moxy/unMock

This is used to remove the existing mock endpoint.

```http request
DELETE /moxy/unMock

{
    "tag": "",
    "method": "",
    "path": ""
}
```

1. **Tag**: This is used to tag the created mock endpoint. It is used to set the context for this endpoint. You can have many instances of the same endpoint.

2. **Method**: This is the request method for this endpoint. It can either be `GET`, `POST`, `PUT ` or `DELETE`.

3. **Path**: This is the request path for this endpoint. Note that the request path should not start with either `/actuator` or `/moxy`.

### Calling any Endpoint

This is used to call any endpoint. If any mock exists for that endpoint, then the mocked response will be returned, otherwise the actual endpoint will be called.

The following 2 mandatory headers should be passed with any such request.

1. **Tag(X-Tag)**: This is the tag used while creating the mock endpoint, to query the mock for the current context.

2. **Actual Base Url(X-Url)**: This is the actual url that will be used to call the actual endpoint, or basically proxy to the actual server, in case no mock exists.

## Example(Using the following actual endpoint - https://www.google.co.in/imghp)

Create a mock endpoint as follows.

```http request
POST /moxy/mock

{
    "tag": "1234",
    "method": "GET",
    "path": "/imghp",
    "responseDelayInMillis": 200,
    "responseStatus": 200,
    "responseBody": {
        "naruto": "rocks"
    }
}
```

Create another mock endpoint with the same request method and path but a different context(tag).

```http request
POST /moxy/mock

{
    "tag": "5678",
    "method": "GET",
    "path": "/imghp",
    "responseDelayInMillis": 200,
    "responseStatus": 200,
    "responseBody": "naruto always rocks"
}
```

Try calling this endpoint with the first tag.

```http request
GET /naruto

X-Tag 1234
X-Url https://www.google.co.in
```

You will get the following response.

```json
{
  "naruto": "rocks"
}
```

Now if you call this with the second tag.

```http request
GET /naruto

X-Tag 5678
X-Url https://www.google.co.in
```

You will get the following response.

```text
naruto always rocks
```

Now try deleting the mock for the first tag.

```http request
DELETE /moxy/unMock

{
    "tag": "1234",
    "method": "GET",
    "path": "/imghp"
}
```

Now again try calling the endpoint with the first tag.

```http request
GET /naruto

X-Tag 1234
X-Url https://www.google.co.in
```

The actual endpoint will be called, and the following response you could see.

```html
<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="en-IN"><head><meta content="Google Images. The most comprehensive image search on the web." name="description"><meta content="text/html; charset=UTF-8" http-equiv="Content-Type"><meta content="/images/branding/googleg/1x/googleg_standard_color_128dp.png" itemprop="image"><title>Google Images</title><script nonce="YYH7DCMo/2WXGV6FU3QJIw==">(function(){window.google={kEI:'Lh_9YKabNIOI-Abk-5CoCg',kEXPI:'0,772215,1,530320,56873,954,755,4350,206,4804,926,1390,383,246,5,1354,5251,1122515,1197765,518,328984,51224,16114,28684,17572,4859,1361,3472,5819,3023,3894,13691,4020,978,13228,3847,4192,6430,1141,7512,5875,234,4282,2778,919,5081,1593,1279,2212,239,291,149,1943,1987,210,4100,108,3406,606,2025,2295,14670,604,2623,2845,7,5599,6755,5096,7876,5037,3407,908,2,941,2614,13142,3,576,1014,1,5444,149,11323,2652,4,1528,2304,1236,5226,577,74,1983,2626,2015,4067,7434,2110,1714,3050,2658,4242,3114,31,13628,2305,639,7079,3772,3494,3269,665,2522,3287,2320,228,992,3102,20,3118,8,906,3,1324,2217,2,8994,5715,1816,281,38,874,5998,12520,2,1394,1525,8,1273,1715,2,3057,723,2,1813,2,1,3,3004,20,1214,3,33,3,5388,3,90,594,784,310,3,435,2379,2422,1274,4578,1576,3,471,578,3,1066,172,3412,2039,101,2,1040,1160,1266,3,3426,2,1712,290,2381,1575,1144,3482,1063,3,123,24,724,3857,758,1447,86,2370,12,2,1458,72,223,2130,408,2,2,5,1858,442,117,394,1996,6,54,70,899,1513,335,87,172,383,51,275,894,158,190,141,427,47,281,85,3,832,280,83,266,1387,523,1258,296,612,1290,332,558,462,123,658,233,635,1609,696,158,434,153,209,1265,103,231,8,316,5609465,99,91,36,220,58,2,69,5996740,2800696,882,444,1,2,80,1,1796,1,9,2,2551,1,748,141,795,563,1,4265,1,1,2,1331,3299,843,2609,155,17,13,72,139,4,2,20,2,169,13,19,46,5,39,96,548,29,2,2,1,2,1,2,2,7,4,1,2,2,2,2,2,2,312,41,513,186,1,1,158,3,2,2,2,2,2,4,2,3,3,236,22,6,5,10,40,2,15,23654791,299865,2867254,1171783,7,2307,277,61,3,2340,74,538,542,5,1701,772339',kBL:'k8M4'};google.sn='imghp';google.kHL='en-IN';})();(function(){
var f=this||self;var h,k=[];function l(a){for(var b;a&&(!a.getAttribute||!(b=a.getAttribute("eid")));)a=a.parentNode;return b||h}function m(a){for(var b=null;a&&(!a.getAttribute||!(b=a.getAttribute("leid")));)a=a.parentNode;return b}
function n(a,b,c,d,g){var e="";c||-1!==b.search("&ei=")||(e="&ei="+l(d),-1===b.search("&lei=")&&(d=m(d))&&(e+="&lei="+d));d="";!c&&f._cshid&&-1===b.search("&cshid=")&&"slh"!==a&&(d="&cshid="+f._cshid);c=c||"/"+(g||"gen_204")+"?atyp=i&ct="+a+"&cad="+b+e+"&zx="+Date.now()+d;/^http:/i.test(c)&&"https:"===window.location.protocol&&(google.ml&&google.ml(Error("a"),!1,{src:c,glmm:1}),c="");return c};h=google.kEI;google.getEI=l;google.getLEI=m;google.ml=function(){return null};google.log=function(a,b,c,d,g){if(c=n(a,b,c,d,g)){a=new Image;var e=k.length;k[e]=a;a.onerror=a.onload=a.onabort=function(){delete k[e]};a.src=c}};google.logUrl=n;}).call(this);(function(){
google.y={};google.sy=[];google.x=function(a,b){if(a)var c=a.id;else{do c=Math.random();while(google.y[c])}google.y[c]=[a,b];return!1};google.sx=function(a){google.sy.push(a)};google.lm=[];google.plm=function(a){google.lm.push.apply(google.lm,a)};google.lq=[];google.load=function(a,b,c){google.lq.push([[a],b,c])};google.loadAll=function(a,b){google.lq.push([a,b])};google.bx=!1;google.lx=function(){};}).call(this);google.f={};(function(){
document.documentElement.addEventListener("submit",function(b){var a;if(a=b.target){var c=a.getAttribute("data-submitfalse");a="1"==c||"q"==c&&!a.elements.q.value?!0:!1}else a=!1;a&&(b.preventDefault(),b.stopPropagation())},!0);document.documentElement.addEventListener("click",function(b){var a;a:{for(a=b.target;a&&a!=document.documentElement;a=a.parentElement)if("A"==a.tagName){a="1"==a.getAttribute("data-nohref");break a}a=!1}a&&b.preventDefault()},!0);}).call(this);</script><style>#gbar,#guser{font-size:13px;padding-top:1px !important;}#gbar{height:22px}#guser{padding-bottom:7px !important;text-align:right}.gbh,.gbd{border-top:1px solid #c9d7f1;font-size:1px}.gbh{height:0;position:absolute;top:24px;width:100%}@media all{.gb1{height:22px;margin-right:.5em;vertical-align:top}#gbar{float:left}}a.gb1,a.gb4{text-decoration:underline !important}a.gb1,a.gb4{color:#00c !important}.gbi .gb4{color:#dd8e27 !important}.gbf .gb4{color:#900 !important}
</style><style>body,td,a,p,.h{font-family:arial,sans-serif}body{margin:0;overflow-y:scroll}#gog{padding:3px 8px 0}td{line-height:.8em}.gac_m td{line-height:17px}form{margin-bottom:20px}.h{color:#1558d6}em{font-weight:bold;font-style:normal}.lst{height:25px;width:496px}.gsfi,.lst{font:18px arial,sans-serif}.gsfs{font:17px arial,sans-serif}.ds{display:inline-box;display:inline-block;margin:3px 0 4px;margin-left:4px}input{font-family:inherit}body{background:#fff;color:#000}a{color:#4b11a8;text-decoration:none}a:hover,a:active{text-decoration:underline}.fl a{color:#1558d6}a:visited{color:#4b11a8}.sblc{padding-top:5px}.sblc a{display:block;margin:2px 0;margin-left:13px;font-size:11px}.lsbb{background:#f8f9fa;border:solid 1px;border-color:#dadce0 #70757a #70757a #dadce0;height:30px}.lsbb{display:block}#WqQANb a{display:inline-block;margin:0 12px}.lsb{background:url(/images/nav_logo229.png) 0 -261px repeat-x;border:none;color:#000;cursor:pointer;height:30px;margin:0;outline:0;font:15px arial,sans-serif;vertical-align:top}.lsb:active{background:#dadce0}.lst:focus{outline:none}.prms{color:#c5221f;font-size:13px}.sshppd{font-size:13px;margin:32px 0 26px}.sshpplo span{color:#c5221f}.sshpplt{margin:15px 0 30px}</style><script nonce="YYH7DCMo/2WXGV6FU3QJIw=="></script></head><body bgcolor="#fff"><script nonce="YYH7DCMo/2WXGV6FU3QJIw==">(function(){var src='/images/nav_logo229.png';var iesg=false;document.body.onload = function(){window.n && window.n();if (document.images){new Image().src=src;}
if (!iesg){document.f&&document.f.q.focus();document.gbqf&&document.gbqf.q.focus();}
}
})();</script><div id="mngb"><div id=gbar><nobr><a class=gb1 href="https://www.google.co.in/webhp?tab=iw">Search</a><b class=gb1>Images</b><a class=gb1 href="https://maps.google.co.in/maps?hl=en&tab=il">Maps</a><a class=gb1 href="https://play.google.com/?hl=en&tab=i8">Play</a><a class=gb1 href="https://www.youtube.com/?gl=IN&tab=i1">YouTube</a><a class=gb1 href="https://news.google.com/?tab=in">News</a><a class=gb1 href="https://mail.google.com/mail/?tab=im">Gmail</a><a class=gb1 href="https://drive.google.com/?tab=io">Drive</a><a class=gb1 style="text-decoration:none" href="https://www.google.co.in/intl/en/about/products?tab=ih"><u>More</u> &raquo;</a></nobr></div><div id=guser width=100%><nobr><span id=gbn class=gbi></span><span id=gbf class=gbf></span><span id=gbe></span><a href="http://www.google.co.in/history/optout?hl=en" class=gb4>Web History</a> |<a  href="/preferences?hl=en" class=gb4>Settings</a> |<a target=_top id=gb_70 href="https://accounts.google.com/ServiceLogin?hl=en&passive=true&continue=https://www.google.co.in/imghp&ec=GAZAAg" class=gb4>Sign in</a></nobr></div><div class=gbh style=left:0></div><div class=gbh style=right:0></div></div><center><br clear="all" id="lgpd"><div id="lga"><div style="padding:28px 0 3px"><div style="height:110px;width:276px;background:url(/intl/en_ALL/images/branding/googlelogo/1x/googlelogo_white_background_color_272x92dp.png) no-repeat" title="Google Images" align="left" id="hplogo"><div style="font-size:16px;font-weight:bold;position:relative;top:70px;color:#1a73e8;right:115px;float:right" nowrap=""><span>images</span></div></div></div><br></div><form action="https://www.google.co.in/search" name="f"><table cellpadding="0" cellspacing="0"><tr valign="top"><td width="25%">&nbsp;</td><td align="center" nowrap=""><input name="tbm" value="isch" type="hidden"><input name="ie" value="ISO-8859-1" type="hidden"><input value="en-IN" name="hl" type="hidden"><input name="source" type="hidden" value="hp"><input name="biw" type="hidden"><input name="bih" type="hidden"><div class="ds" style="height:32px;margin:4px 0"><input class="lst" style="margin:0;padding:5px 8px 0 6px;vertical-align:top;color:#000" autocomplete="off" value="" title="Search Images" maxlength="2048" name="q" size="57"></div><span class="ds"><span class="lsbb"><input class="lsb" value="Search Images" name="btnG" type="submit"></span></span></td><td class="fl sblc" align="left" valign="middle" nowrap="" width="25%"><a href="/advanced_image_search?hl=en-IN&amp;authuser=0">Advanced&nbsp;Image&nbsp;Search</a></td></tr></table><input id="gbv" name="gbv" type="hidden" value="1"><script nonce="YYH7DCMo/2WXGV6FU3QJIw==">(function(){
var a,b="1";if(document&&document.getElementById)if("undefined"!=typeof XMLHttpRequest)b="2";else if("undefined"!=typeof ActiveXObject){var c,d,e=["MSXML2.XMLHTTP.6.0","MSXML2.XMLHTTP.3.0","MSXML2.XMLHTTP","Microsoft.XMLHTTP"];for(c=0;d=e[c++];)try{new ActiveXObject(d),b="2"}catch(h){}}a=b;if("2"==a&&-1==location.search.indexOf("&gbv=2")){var f=google.gbvu,g=document.getElementById("gbv");g&&(g.value=a);f&&window.setTimeout(function(){location.href=f},0)};}).call(this);</script></form><div id="gac_scont"></div><div style="font-size:83%;min-height:3.5em"><br></div><span id="footer"><div style="font-size:10pt"><div style="margin:19px auto;text-align:center" id="WqQANb"><a href="/intl/en/ads/">Advertising Programs</a><a href="http://www.google.co.in/services/">Business Solutions</a><a href="/intl/en/about.html">About Google</a></div></div><p style="font-size:8pt;color:#70757a">&copy; 2021 - <a href="/intl/en/policies/privacy/">Privacy</a> -<a href="/intl/en/policies/terms/">Terms</a></p></span></center><script nonce="YYH7DCMo/2WXGV6FU3QJIw==">(function(){window.google.cdo={height:757,width:1440};(function(){
var a=window.innerWidth,b=window.innerHeight;if(!a||!b){var c=window.document,d="CSS1Compat"==c.compatMode?c.documentElement:c.body;a=d.clientWidth;b=d.clientHeight}a&&b&&(a!=google.cdo.width||b!=google.cdo.height)&&google.log("","","/client_204?&atyp=i&biw="+a+"&bih="+b+"&ei="+google.kEI);}).call(this);})();</script><script nonce="YYH7DCMo/2WXGV6FU3QJIw==">(function(){google.xjs={ck:'',cs:'',excm:[],pml:false};})();</script><script nonce="YYH7DCMo/2WXGV6FU3QJIw==">(function(){var u='/xjs/_/js/k\x3dxjs.hp.en.9gLGBHDOM4E.O/m\x3dsb_he,d/am\x3dAPgEWA/d\x3d1/ed\x3d1/rs\x3dACT90oEXm9MrN4B8xXQY6N86St4ZoOKYHQ';
var e=this||self,f=function(a){return a};var g;var l=function(a,b){this.g=b===h?a:""};l.prototype.toString=function(){return this.g+""};var h={};function m(){var a=u;google.lx=function(){n(a);google.lx=function(){}};google.bx||google.lx()}
function n(a){google.timers&&google.timers.load&&google.tick&&google.tick("load","xjsls");var b=document;var c="SCRIPT";"application/xhtml+xml"===b.contentType&&(c=c.toLowerCase());c=b.createElement(c);if(void 0===g){b=null;var k=e.trustedTypes;if(k&&k.createPolicy){try{b=k.createPolicy("goog#html",{createHTML:f,createScript:f,createScriptURL:f})}catch(p){e.console&&e.console.error(p.message)}g=b}else g=b}a=(b=g)?b.createScriptURL(a):a;a=new l(a,h);c.src=a instanceof l&&a.constructor===l?a.g:"type_error:TrustedResourceUrl";var d;a=(c.ownerDocument&&c.ownerDocument.defaultView||window).document;(d=(a=null===(d=a.querySelector)||void 0===d?void 0:d.call(a,"script[nonce]"))?a.nonce||a.getAttribute("nonce")||"":"")&&c.setAttribute("nonce",d);document.body.appendChild(c);google.psa=!0};setTimeout(function(){m()},0);})();(function(){window.google.xjsu='/xjs/_/js/k\x3dxjs.hp.en.9gLGBHDOM4E.O/m\x3dsb_he,d/am\x3dAPgEWA/d\x3d1/ed\x3d1/rs\x3dACT90oEXm9MrN4B8xXQY6N86St4ZoOKYHQ';})();function _DumpException(e){throw e;}
function _F_installCss(c){}
(function(){google.jl={attn:false,blt:'none',chnk:0,dw:false,emtn:0,end:0,ine:false,lls:'default',pdt:0,rep:0,sif:false,snet:true,strt:0,ubm:false,uwp:true};})();(function(){var pmc='{\x22d\x22:{},\x22sb_he\x22:{\x22agen\x22:true,\x22cgen\x22:true,\x22client\x22:\x22img\x22,\x22dh\x22:true,\x22dhqt\x22:true,\x22ds\x22:\x22i\x22,\x22ffql\x22:\x22en\x22,\x22host\x22:\x22google.co.in\x22,\x22isbh\x22:28,\x22jsonp\x22:true,\x22msgs\x22:{\x22cibl\x22:\x22Clear Search\x22,\x22dym\x22:\x22Did you mean:\x22,\x22lcky\x22:\x22I\\u0026#39;m Feeling Lucky\x22,\x22lml\x22:\x22Learn more\x22,\x22oskt\x22:\x22Input tools\x22,\x22psrc\x22:\x22This search was removed from your \\u003Ca href\x3d\\\x22/history\\\x22\\u003EWeb History\\u003C/a\\u003E\x22,\x22psrl\x22:\x22Remove\x22,\x22sbit\x22:\x22Search by image\x22,\x22srch\x22:\x22Google Search\x22},\x22ovr\x22:{},\x22pq\x22:\x22\x22,\x22sbas\x22:\x220 3px 8px 0 rgba(0,0,0,0.2),0 0 0 1px rgba(0,0,0,0.08)\x22,\x22sbpl\x22:16,\x22sbpr\x22:16,\x22scd\x22:10,\x22stok\x22:\x22qb6u1LJ3VhKtpclmepiiAm7jA9E\x22,\x22uhde\x22:false}}';google.pmc=JSON.parse(pmc);})();</script></body></html>
```


