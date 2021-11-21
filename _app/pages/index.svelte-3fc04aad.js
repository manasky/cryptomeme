import{S as Te,i as Ie,s as Se,J as tt,e as p,k,t as V,E as O,c as v,a as _,n as y,d as o,g as x,F as U,b as i,O as Ye,f as b,G as u,h as ce,L as lt,M as rt,N as st,x as q,u as te,D as fe,P as it,A as qe,Q as at,R as nt,j as Ce,K as ot,m as Be,o as Ve,w as Le,v as xe,r as ze,l as He,H as Je}from"../chunks/vendor-02e82060.js";const ct=[["y",31536e3],["mo",2592e3],["d",86400],["h",3600],["m",60],["s",1]],ft=a=>{for(let[e,r]of ct){const l=Math.floor(a/r);if(l>=1)return{interval:l,epoch:e}}},Ke=a=>{const e=Math.floor((new Date-new Date(a))/1e3),{interval:r,epoch:l}=ft(e);return`${r}${l}`};function ut(a){let e,r,l;return{c(){e=O("svg"),r=O("circle"),l=O("path"),this.h()},l(t){e=U(t,"svg",{class:!0,xmlns:!0,fill:!0,viewBox:!0});var s=_(e);r=U(s,"circle",{class:!0,cx:!0,cy:!0,r:!0,stroke:!0,"stroke-width":!0}),_(r).forEach(o),l=U(s,"path",{class:!0,fill:!0,d:!0}),_(l).forEach(o),s.forEach(o),this.h()},h(){i(r,"class","opacity-25"),i(r,"cx","12"),i(r,"cy","12"),i(r,"r","10"),i(r,"stroke","currentColor"),i(r,"stroke-width","4"),i(l,"class","opacity-75"),i(l,"fill","currentColor"),i(l,"d","M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"),i(e,"class","w-4 h-4 animate-spin text-base-content"),i(e,"xmlns","http://www.w3.org/2000/svg"),i(e,"fill","none"),i(e,"viewBox","0 0 24 24")},m(t,s){b(t,e,s),u(e,r),u(e,l)},p:fe,d(t){t&&o(e)}}}function ht(a){let e=a[0].name+"",r;return{c(){r=V(e)},l(l){r=x(l,e)},m(l,t){b(l,r,t)},p(l,t){t&1&&e!==(e=l[0].name+"")&&ce(r,e)},d(l){l&&o(r)}}}function mt(a){let e,r,l;return{c(){e=O("svg"),r=O("circle"),l=O("path"),this.h()},l(t){e=U(t,"svg",{class:!0,xmlns:!0,fill:!0,viewBox:!0});var s=_(e);r=U(s,"circle",{class:!0,cx:!0,cy:!0,r:!0,stroke:!0,"stroke-width":!0}),_(r).forEach(o),l=U(s,"path",{class:!0,fill:!0,d:!0}),_(l).forEach(o),s.forEach(o),this.h()},h(){i(r,"class","opacity-25"),i(r,"cx","12"),i(r,"cy","12"),i(r,"r","10"),i(r,"stroke","currentColor"),i(r,"stroke-width","4"),i(l,"class","opacity-75"),i(l,"fill","currentColor"),i(l,"d","M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"),i(e,"class","w-4 h-4 animate-spin text-base-content"),i(e,"xmlns","http://www.w3.org/2000/svg"),i(e,"fill","none"),i(e,"viewBox","0 0 24 24")},m(t,s){b(t,e,s),u(e,r),u(e,l)},p:fe,d(t){t&&o(e)}}}function _t(a){let e,r=a[1](a[0].current_price)+"",l;return{c(){e=V("$"),l=V(r)},l(t){e=x(t,"$"),l=x(t,r)},m(t,s){b(t,e,s),b(t,l,s)},p(t,s){s&1&&r!==(r=t[1](t[0].current_price)+"")&&ce(l,r)},d(t){t&&o(e),t&&o(l)}}}function dt(a){let e,r,l,t,s=Number(Math.abs(a[0].price_change_percentage_24h)).toFixed(2)+"",n,h,m;return{c(){e=O("svg"),r=O("path"),l=k(),t=p("span"),n=V(s),h=V("%"),this.h()},l(f){e=U(f,"svg",{class:!0,xmlns:!0,viewBox:!0});var g=_(e);r=U(g,"path",{d:!0}),_(r).forEach(o),g.forEach(o),l=y(f),t=v(f,"SPAN",{class:!0,"data-tip":!0});var w=_(t);n=x(w,s),h=x(w,"%"),w.forEach(o),this.h()},h(){i(r,"d","M17,6a1,1,0,0,0-1,1v7.59L7.71,6.29A1,1,0,0,0,6.29,7.71L14.59,16H7a1,1,0,0,0,0,2H17a1,1,0,0,0,.38-.08,1,1,0,0,0,.54-.54A1,1,0,0,0,18,17V7A1,1,0,0,0,17,6Z"),i(e,"class","w-4 h-4 fill-current text-error"),i(e,"xmlns","http://www.w3.org/2000/svg"),i(e,"viewBox","0 0 24 24"),i(t,"class","z-[2] font-mono text-sm text-error tooltip tooltip-right"),i(t,"data-tip",m="-$"+Number(Math.abs(a[0].price_change_24h))+" in 24 hours")},m(f,g){b(f,e,g),u(e,r),b(f,l,g),b(f,t,g),u(t,n),u(t,h)},p(f,g){g&1&&s!==(s=Number(Math.abs(f[0].price_change_percentage_24h)).toFixed(2)+"")&&ce(n,s),g&1&&m!==(m="-$"+Number(Math.abs(f[0].price_change_24h))+" in 24 hours")&&i(t,"data-tip",m)},d(f){f&&o(e),f&&o(l),f&&o(t)}}}function pt(a){let e,r,l,t,s=Number(Math.abs(a[0].price_change_percentage_24h)).toFixed(2)+"",n,h,m;return{c(){e=O("svg"),r=O("path"),l=k(),t=p("span"),n=V(s),h=V("%"),this.h()},l(f){e=U(f,"svg",{class:!0,xmlns:!0,viewBox:!0});var g=_(e);r=U(g,"path",{d:!0}),_(r).forEach(o),g.forEach(o),l=y(f),t=v(f,"SPAN",{class:!0,"data-tip":!0});var w=_(t);n=x(w,s),h=x(w,"%"),w.forEach(o),this.h()},h(){i(r,"d","M17.92,6.62a1,1,0,0,0-.54-.54A1,1,0,0,0,17,6H7A1,1,0,0,0,7,8h7.59l-8.3,8.29a1,1,0,0,0,0,1.42,1,1,0,0,0,1.42,0L16,9.41V17a1,1,0,0,0,2,0V7A1,1,0,0,0,17.92,6.62Z"),i(e,"class","w-4 h-4 fill-current text-success"),i(e,"xmlns","http://www.w3.org/2000/svg"),i(e,"viewBox","0 0 24 24"),i(t,"class","z-[2] font-mono text-sm text-success tooltip tooltip-right"),i(t,"data-tip",m="+$"+Number(Math.abs(a[0].price_change_24h))+" in 24 hours")},m(f,g){b(f,e,g),u(e,r),b(f,l,g),b(f,t,g),u(t,n),u(t,h)},p(f,g){g&1&&s!==(s=Number(Math.abs(f[0].price_change_percentage_24h)).toFixed(2)+"")&&ce(n,s),g&1&&m!==(m="+$"+Number(Math.abs(f[0].price_change_24h))+" in 24 hours")&&i(t,"data-tip",m)},d(f){f&&o(e),f&&o(l),f&&o(t)}}}function vt(a){let e,r,l,t,s,n,h,m,f,g,w,T=a[0].symbol+"",L,I,N,M,z,S,C,j,c,d,Q,X,J,se,le,me,Z=a[1](a[0].high_24h)+"",ue,_e,G,H,F,he,ie,ae,B=a[1](a[0].low_24h)+"",de,ee,ne,E;function W(A,D){return A[0].name?ht:ut}let re=W(a),$=re(a);function Ae(A,D){return A[0].current_price?_t:mt}let K=Ae(a),Y=K(a);function pe(A,D){if(A[0].price_change_percentage_24h>0)return pt;if(A[0].price_change_percentage_24h<0)return dt}let oe=pe(a),P=oe&&oe(a);const be=a[3].default,R=tt(be,a,a[2],null);return{c(){e=p("tr"),r=p("th"),l=p("div"),t=p("img"),h=k(),m=p("div"),f=p("div"),$.c(),g=k(),w=p("div"),L=V(T),I=k(),N=p("td"),M=p("div"),z=p("div"),Y.c(),S=k(),C=p("div"),P&&P.c(),j=k(),c=p("td"),d=p("div"),Q=p("div"),X=O("svg"),J=O("path"),se=k(),le=p("span"),me=V("Max: $"),ue=V(Z),_e=k(),G=p("div"),H=O("svg"),F=O("path"),he=k(),ie=p("span"),ae=V("Min: $"),de=V(B),ee=k(),ne=p("td"),R&&R.c(),this.h()},l(A){e=v(A,"TR",{class:!0});var D=_(e);r=v(D,"TH",{class:!0});var ve=_(r);l=v(ve,"DIV",{class:!0});var ge=_(l);t=v(ge,"IMG",{class:!0,src:!0,alt:!0,width:!0,height:!0}),h=y(ge),m=v(ge,"DIV",{class:!0});var we=_(m);f=v(we,"DIV",{class:!0});var Ee=_(f);$.l(Ee),Ee.forEach(o),g=y(we),w=v(we,"DIV",{class:!0});var Oe=_(w);L=x(Oe,T),Oe.forEach(o),we.forEach(o),ge.forEach(o),ve.forEach(o),I=y(D),N=v(D,"TD",{});var Ue=_(N);M=v(Ue,"DIV",{class:!0});var ke=_(M);z=v(ke,"DIV",{class:!0});var Fe=_(z);Y.l(Fe),Fe.forEach(o),S=y(ke),C=v(ke,"DIV",{class:!0});var Re=_(C);P&&P.l(Re),Re.forEach(o),ke.forEach(o),Ue.forEach(o),j=y(D),c=v(D,"TD",{});var $e=_(c);d=v($e,"DIV",{class:!0});var ye=_(d);Q=v(ye,"DIV",{class:!0});var Me=_(Q);X=U(Me,"svg",{class:!0,xmlns:!0,viewBox:!0});var je=_(X);J=U(je,"path",{d:!0}),_(J).forEach(o),je.forEach(o),se=y(Me),le=v(Me,"SPAN",{class:!0});var Ne=_(le);me=x(Ne,"Max: $"),ue=x(Ne,Z),Ne.forEach(o),Me.forEach(o),_e=y(ye),G=v(ye,"DIV",{class:!0});var De=_(G);H=U(De,"svg",{class:!0,xmlns:!0,viewBox:!0});var Ge=_(H);F=U(Ge,"path",{d:!0}),_(F).forEach(o),Ge.forEach(o),he=y(De),ie=v(De,"SPAN",{class:!0});var Pe=_(ie);ae=x(Pe,"Min: $"),de=x(Pe,B),Pe.forEach(o),De.forEach(o),ye.forEach(o),$e.forEach(o),ee=y(D),ne=v(D,"TD",{class:!0});var We=_(ne);R&&R.l(We),We.forEach(o),D.forEach(o),this.h()},h(){i(t,"class","w-8 h-8 rounded"),Ye(t.src,s=a[0].image)||i(t,"src",s),i(t,"alt",n=a[0].name),i(t,"width","32"),i(t,"height","32"),i(f,"class","font-semibold"),i(w,"class","font-mono text-xs text-opacity-70 font-light uppercase text-base-content"),i(m,"class","flex flex-col pr-4"),i(l,"class","flex items-center gap-3"),i(r,"class","flex-shrink-0 rounded-l-lg bg-primary"),i(z,"class","font-mono text-lg font-bold"),i(C,"class","flex items-center gap-1"),i(M,"class","flex flex-col"),i(J,"d","M1024 512l1024 1024H0L1024 512z"),i(X,"class","w-2 h-2 fill-current opacity-30"),i(X,"xmlns","http://www.w3.org/2000/svg"),i(X,"viewBox","0 0 2048 2048"),i(le,"class","opacity-70"),i(Q,"class","flex items-center gap-1 text-xs"),i(F,"d","M0 512h2048L1024 1536 0 512z"),i(H,"class","w-2 h-2 fill-current opacity-30"),i(H,"xmlns","http://www.w3.org/2000/svg"),i(H,"viewBox","0 0 2048 2048"),i(ie,"class","opacity-70"),i(G,"class","flex items-center gap-1 text-xs"),i(d,"class","flex flex-col"),i(ne,"class","flex-shrink-0 rounded-r-lg"),i(e,"class","hover")},m(A,D){b(A,e,D),u(e,r),u(r,l),u(l,t),u(l,h),u(l,m),u(m,f),$.m(f,null),u(m,g),u(m,w),u(w,L),u(e,I),u(e,N),u(N,M),u(M,z),Y.m(z,null),u(M,S),u(M,C),P&&P.m(C,null),u(e,j),u(e,c),u(c,d),u(d,Q),u(Q,X),u(X,J),u(Q,se),u(Q,le),u(le,me),u(le,ue),u(d,_e),u(d,G),u(G,H),u(H,F),u(G,he),u(G,ie),u(ie,ae),u(ie,de),u(e,ee),u(e,ne),R&&R.m(ne,null),E=!0},p(A,[D]){(!E||D&1&&!Ye(t.src,s=A[0].image))&&i(t,"src",s),(!E||D&1&&n!==(n=A[0].name))&&i(t,"alt",n),re===(re=W(A))&&$?$.p(A,D):($.d(1),$=re(A),$&&($.c(),$.m(f,null))),(!E||D&1)&&T!==(T=A[0].symbol+"")&&ce(L,T),K===(K=Ae(A))&&Y?Y.p(A,D):(Y.d(1),Y=K(A),Y&&(Y.c(),Y.m(z,null))),oe===(oe=pe(A))&&P?P.p(A,D):(P&&P.d(1),P=oe&&oe(A),P&&(P.c(),P.m(C,null))),(!E||D&1)&&Z!==(Z=A[1](A[0].high_24h)+"")&&ce(ue,Z),(!E||D&1)&&B!==(B=A[1](A[0].low_24h)+"")&&ce(de,B),R&&R.p&&(!E||D&4)&&lt(R,be,A,A[2],E?st(be,A[2],D,null):rt(A[2]),null)},i(A){E||(q(R,A),E=!0)},o(A){te(R,A),E=!1},d(A){A&&o(e),$.d(),Y.d(),P&&P.d(),R&&R.d(A)}}}function gt(a,e,r){let{$$slots:l={},$$scope:t}=e,{coin:s={name:"",symbol:"",current_price:"",price_change_24h:"",price_change_percentage_24h:"",high_24h:"",low_24h:"",image:"data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==",chart:"data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=="}}=e,n=h=>h>1?h.toLocaleString("en-US",{maximumFractionDigits:2}):h.toLocaleString("en-US",{maximumFractionDigits:10});return a.$$set=h=>{"coin"in h&&r(0,s=h.coin),"$$scope"in h&&r(2,t=h.$$scope)},[s,n,t,l]}class Qe extends Te{constructor(e){super();Ie(this,e,gt,vt,Se,{coin:0})}}const wt=(a,e,r,l,t,s)=>{let n=e.split(" "),h="";for(let m=0;m<n.length;m++){let f=h+n[m]+" ";a.measureText(f).width>t&&m>0?(a.fillText(h,r,l),h=n[m]+" ",l+=s):h=f}a.fillText(h,r,l)},bt=({canvas:a,ctx:e,path:r,creditPrice:l="",creditDate:t="",meme_caption:s=""})=>{let n=new Image;n.src=r,n.onload=function(){e.fillStyle="#fff",e.fillRect(0,0,a.width,a.height);let h=a.width*n.height/n.width,m=a.height-h;e.drawImage(n,0,m,a.width,h);let f;s.length>300?f=12:s.length>200?f=18:s.length>180?f=22:s.length>150?f=24:s.length>100?f=32:s.length>40?f=40:s.length>30?f=54:s.length>10?f=60:f=68;let g=a.width-40,w=f+f*.4,T=f+20,L=(a.width-g)/2;e.font="Bold "+f+"px 'Helvetica'",e.fillStyle="black",wt(e,s,L,T,g,w);let I="CryptoMeme.WTF",N=({ctx:M,text:z,textX:S,textY:C})=>{M.font="10px 'Arial'",M.strokeStyle="rgba(0,0,0,.3)",M.lineWidth=3,M.strokeText(z,S,C),M.fillStyle="rgba(255,255,255,0.6)",M.fillText(z,S,C)};N({ctx:e,text:t,textX:10,textY:a.height-20}),N({ctx:e,text:l,textX:9,textY:a.height-10}),N({ctx:e,text:I,textX:a.width-90,textY:a.height-10})}};function At(a){let e,r,l;return{c(){e=p("canvas"),this.h()},l(t){e=v(t,"CANVAS",{class:!0,width:!0,height:!0,id:!0}),_(e).forEach(o),this.h()},h(){i(e,"class","w-52 -my-3 -mr-4 border border-base-content border-opacity-5 cursor-pointer rounded"),i(e,"width","700"),i(e,"height","600"),i(e,"id",a[0])},m(t,s){b(t,e,s),a[7](e),r||(l=it(e,"click",function(){at(a[2]())&&a[2]().apply(this,arguments)}),r=!0)},p(t,[s]){a=t,s&1&&i(e,"id",a[0])},i:fe,o:fe,d(t){t&&o(e),a[7](null),r=!1,l()}}}function Et(a,e,r){let{id:l}=e,{meme:t}=e,{meme_caption:s}=e,{creditPrice:n=""}=e,{creditDate:h=""}=e,m,f;qe(()=>{let w=m.getContext("2d");bt({canvas:m,ctx:w,path:t,meme_caption:s,creditPrice:n,creditDate:h}),r(2,f=()=>{let T=document.createElement("a");T.download=l+"-meme.jpg",T.href=m.toDataURL(),T.click()})});function g(w){nt[w?"unshift":"push"](()=>{m=w,r(1,m)})}return a.$$set=w=>{"id"in w&&r(0,l=w.id),"meme"in w&&r(3,t=w.meme),"meme_caption"in w&&r(4,s=w.meme_caption),"creditPrice"in w&&r(5,n=w.creditPrice),"creditDate"in w&&r(6,h=w.creditDate)},[l,m,f,t,s,n,h,g]}class kt extends Te{constructor(e){super();Ie(this,e,Et,At,Se,{id:0,meme:3,meme_caption:4,creditPrice:5,creditDate:6})}}function yt(a){let e,r,l,t,s,n,h,m,f,g,w,T,L,I,N,M,z,S,C,j;return{c(){e=p("title"),r=V(a[0]),l=k(),t=p("meta"),s=k(),n=p("meta"),h=k(),m=p("meta"),f=k(),g=p("meta"),w=k(),T=p("meta"),L=k(),I=p("meta"),N=k(),M=p("meta"),z=k(),S=p("meta"),C=k(),j=p("meta"),this.h()},l(c){e=v(c,"TITLE",{});var d=_(e);r=x(d,a[0]),d.forEach(o),l=y(c),t=v(c,"META",{name:!0,content:!0}),s=y(c),n=v(c,"META",{name:!0,content:!0}),h=y(c),m=v(c,"META",{name:!0,content:!0}),f=y(c),g=v(c,"META",{name:!0,content:!0}),w=y(c),T=v(c,"META",{name:!0,content:!0}),L=y(c),I=v(c,"META",{name:!0,content:!0}),N=y(c),M=v(c,"META",{property:!0,content:!0}),z=y(c),S=v(c,"META",{property:!0,content:!0}),C=y(c),j=v(c,"META",{property:!0,content:!0}),this.h()},h(){i(t,"name","description"),i(t,"content",a[1]),i(n,"name","twitter:card"),i(n,"content","summary"),i(m,"name","twitter:title"),i(m,"content",a[0]),i(g,"name","twitter:description"),i(g,"content",a[1]),i(T,"name","twitter:image"),i(T,"content",a[2]),i(I,"name","twitter:image:alt"),i(I,"content",a[0]),i(M,"property","og:title"),i(M,"content",a[0]),i(S,"property","og:description"),i(S,"content",a[1]),i(j,"property","og:image"),i(j,"content",a[2])},m(c,d){b(c,e,d),u(e,r),b(c,l,d),b(c,t,d),b(c,s,d),b(c,n,d),b(c,h,d),b(c,m,d),b(c,f,d),b(c,g,d),b(c,w,d),b(c,T,d),b(c,L,d),b(c,I,d),b(c,N,d),b(c,M,d),b(c,z,d),b(c,S,d),b(c,C,d),b(c,j,d)},p(c,[d]){d&1&&ce(r,c[0]),d&2&&i(t,"content",c[1]),d&1&&i(m,"content",c[0]),d&2&&i(g,"content",c[1]),d&4&&i(T,"content",c[2]),d&1&&i(I,"content",c[0]),d&1&&i(M,"content",c[0]),d&2&&i(S,"content",c[1]),d&4&&i(j,"content",c[2])},i:fe,o:fe,d(c){c&&o(e),c&&o(l),c&&o(t),c&&o(s),c&&o(n),c&&o(h),c&&o(m),c&&o(f),c&&o(g),c&&o(w),c&&o(T),c&&o(L),c&&o(I),c&&o(N),c&&o(M),c&&o(z),c&&o(S),c&&o(C),c&&o(j)}}}function Mt(a,e,r){let{title:l}=e,{desc:t}=e,{img:s}=e;return a.$$set=n=>{"title"in n&&r(0,l=n.title),"desc"in n&&r(1,t=n.desc),"img"in n&&r(2,s=n.img)},[l,t,s]}class Dt extends Te{constructor(e){super();Ie(this,e,Mt,yt,Se,{title:0,desc:1,img:2})}}function Tt(a,e,r){const l=a.slice();return l[4]=e[r],l}function Xe(a,e,r){const l=a.slice();return l[4]=e[r],l}function It(a){let e;function r(s,n){return s[1]?Bt:Ct}let l=r(a),t=l(a);return{c(){e=p("span"),t.c(),this.h()},l(s){e=v(s,"SPAN",{class:!0});var n=_(e);t.l(n),n.forEach(o),this.h()},h(){i(e,"class","text-xs normal-case text-warning")},m(s,n){b(s,e,n),t.m(e,null)},p(s,n){l===(l=r(s))&&t?t.p(s,n):(t.d(1),t=l(s),t&&(t.c(),t.m(e,null)))},d(s){s&&o(e),t.d()}}}function St(a){let e,r,l;return{c(){e=O("svg"),r=O("circle"),l=O("path"),this.h()},l(t){e=U(t,"svg",{class:!0,xmlns:!0,fill:!0,viewBox:!0});var s=_(e);r=U(s,"circle",{class:!0,cx:!0,cy:!0,r:!0,stroke:!0,"stroke-width":!0}),_(r).forEach(o),l=U(s,"path",{class:!0,fill:!0,d:!0}),_(l).forEach(o),s.forEach(o),this.h()},h(){i(r,"class","opacity-25"),i(r,"cx","12"),i(r,"cy","12"),i(r,"r","10"),i(r,"stroke","currentColor"),i(r,"stroke-width","4"),i(l,"class","opacity-75"),i(l,"fill","currentColor"),i(l,"d","M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"),i(e,"class","w-4 h-4 animate-spin text-base-content"),i(e,"xmlns","http://www.w3.org/2000/svg"),i(e,"fill","none"),i(e,"viewBox","0 0 24 24")},m(t,s){b(t,e,s),u(e,r),u(e,l)},p:fe,d(t){t&&o(e)}}}function Ct(a){let e;return{c(){e=V("Cannot connect to server")},l(r){e=x(r,"Cannot connect to server")},m(r,l){b(r,e,l)},p:fe,d(r){r&&o(e)}}}function Bt(a){let e,r=Ke(a[1])+"",l;return{c(){e=V("Last update: "),l=V(r)},l(t){e=x(t,"Last update: "),l=x(t,r)},m(t,s){b(t,e,s),b(t,l,s)},p(t,s){s&2&&r!==(r=Ke(t[1])+"")&&ce(l,r)},d(t){t&&o(e),t&&o(l)}}}function Vt(a){let e,r,l=Array(10),t=[];for(let s=0;s<l.length;s+=1)t[s]=Ht(Tt(a,l,s));return{c(){for(let s=0;s<t.length;s+=1)t[s].c();e=He()},l(s){for(let n=0;n<t.length;n+=1)t[n].l(s);e=He()},m(s,n){for(let h=0;h<t.length;h+=1)t[h].m(s,n);b(s,e,n),r=!0},p:fe,i(s){if(!r){for(let n=0;n<l.length;n+=1)q(t[n]);r=!0}},o(s){t=t.filter(Boolean);for(let n=0;n<t.length;n+=1)te(t[n]);r=!1},d(s){Je(t,s),s&&o(e)}}}function xt(a){let e,r,l=a[0].slice(0,10),t=[];for(let n=0;n<l.length;n+=1)t[n]=et(Xe(a,l,n));const s=n=>te(t[n],1,1,()=>{t[n]=null});return{c(){for(let n=0;n<t.length;n+=1)t[n].c();e=He()},l(n){for(let h=0;h<t.length;h+=1)t[h].l(n);e=He()},m(n,h){for(let m=0;m<t.length;m+=1)t[m].m(n,h);b(n,e,h),r=!0},p(n,h){if(h&1){l=n[0].slice(0,10);let m;for(m=0;m<l.length;m+=1){const f=Xe(n,l,m);t[m]?(t[m].p(f,h),q(t[m],1)):(t[m]=et(f),t[m].c(),q(t[m],1),t[m].m(e.parentNode,e))}for(ze(),m=l.length;m<t.length;m+=1)s(m);Le()}},i(n){if(!r){for(let h=0;h<l.length;h+=1)q(t[h]);r=!0}},o(n){t=t.filter(Boolean);for(let h=0;h<t.length;h+=1)te(t[h]);r=!1},d(n){Je(t,n),n&&o(e)}}}function Ht(a){let e,r;return e=new Qe({}),{c(){Ce(e.$$.fragment)},l(l){Be(e.$$.fragment,l)},m(l,t){Ve(e,l,t),r=!0},i(l){r||(q(e.$$.fragment,l),r=!0)},o(l){te(e.$$.fragment,l),r=!1},d(l){xe(e,l)}}}function Ze(a){let e,r;return e=new kt({props:{id:a[4].symbol,meme:a[4].meme,meme_caption:a[4].meme_caption,creditDate:"",creditPrice:""+(a[4].symbol.toUpperCase()+" price on "+new Date(Date.now()).toLocaleString("en-US",{year:"numeric",month:"short",day:"2-digit"})+": $"+a[4].current_price+" ("+((a[4].price_change_24h<0?"\u25BC":"\u25B2")+"$"+Number(Math.abs(a[4].price_change_24h)))+")")}}),{c(){Ce(e.$$.fragment)},l(l){Be(e.$$.fragment,l)},m(l,t){Ve(e,l,t),r=!0},p(l,t){const s={};t&1&&(s.id=l[4].symbol),t&1&&(s.meme=l[4].meme),t&1&&(s.meme_caption=l[4].meme_caption),t&1&&(s.creditPrice=""+(l[4].symbol.toUpperCase()+" price on "+new Date(Date.now()).toLocaleString("en-US",{year:"numeric",month:"short",day:"2-digit"})+": $"+l[4].current_price+" ("+((l[4].price_change_24h<0?"\u25BC":"\u25B2")+"$"+Number(Math.abs(l[4].price_change_24h)))+")")),e.$set(s)},i(l){r||(q(e.$$.fragment,l),r=!0)},o(l){te(e.$$.fragment,l),r=!1},d(l){xe(e,l)}}}function Nt(a){let e,r,l=a[4].meme&&Ze(a);return{c(){l&&l.c(),e=k()},l(t){l&&l.l(t),e=y(t)},m(t,s){l&&l.m(t,s),b(t,e,s),r=!0},p(t,s){t[4].meme?l?(l.p(t,s),s&1&&q(l,1)):(l=Ze(t),l.c(),q(l,1),l.m(e.parentNode,e)):l&&(ze(),te(l,1,1,()=>{l=null}),Le())},i(t){r||(q(l),r=!0)},o(t){te(l),r=!1},d(t){l&&l.d(t),t&&o(e)}}}function et(a){let e,r;return e=new Qe({props:{coin:a[4],$$slots:{default:[Nt]},$$scope:{ctx:a}}}),{c(){Ce(e.$$.fragment)},l(l){Be(e.$$.fragment,l)},m(l,t){Ve(e,l,t),r=!0},p(l,t){const s={};t&1&&(s.coin=l[4]),t&513&&(s.$$scope={dirty:t,ctx:l}),e.$set(s)},i(l){r||(q(e.$$.fragment,l),r=!0)},o(l){te(e.$$.fragment,l),r=!1},d(l){xe(e,l)}}}function Pt(a){let e,r,l,t,s,n,h,m,f,g,w,T,L,I,N,M,z,S,C,j,c,d,Q,X,J,se,le,me,Z,ue,_e,G,H,F,he;e=new Dt({props:{title:"CryptoMeme.WTF - Crypto Price + Memes",desc:"Crypto price memes - Bitcoin meme, Ethereum meme, crypto meme",img:"/card.jpg"}});function ie(E,W){if(E[2]=="loading")return St;if(E[2]=="error")return It}let ae=ie(a),B=ae&&ae(a);const de=[xt,Vt],ee=[];function ne(E,W){return E[0]?0:1}return H=ne(a),F=ee[H]=de[H](a),{c(){Ce(e.$$.fragment),r=k(),l=p("div"),t=p("table"),s=p("thead"),n=p("tr"),h=p("th"),B&&B.c(),m=k(),f=p("th"),g=p("div"),w=V("Current Price"),T=k(),L=p("div"),I=O("svg"),N=O("path"),M=V(`
            Price change in 24 hours`),z=k(),S=p("th"),C=p("div"),j=V("Min/Max price"),c=k(),d=p("div"),Q=V("In 24 hours"),X=k(),J=p("th"),se=p("div"),le=V("Current price meme"),me=k(),Z=p("div"),ue=V("Generated hourly"),_e=k(),G=p("tbody"),F.c(),this.h()},l(E){const W=ot('[data-svelte="svelte-myrbta"]',document.head);Be(e.$$.fragment,W),W.forEach(o),r=y(E),l=v(E,"DIV",{class:!0});var re=_(l);t=v(re,"TABLE",{class:!0});var $=_(t);s=v($,"THEAD",{});var Ae=_(s);n=v(Ae,"TR",{});var K=_(n);h=v(K,"TH",{});var Y=_(h);B&&B.l(Y),Y.forEach(o),m=y(K),f=v(K,"TH",{});var pe=_(f);g=v(pe,"DIV",{class:!0});var oe=_(g);w=x(oe,"Current Price"),oe.forEach(o),T=y(pe),L=v(pe,"DIV",{class:!0});var P=_(L);I=U(P,"svg",{class:!0,xmlns:!0,viewBox:!0});var be=_(I);N=U(be,"path",{d:!0}),_(N).forEach(o),be.forEach(o),M=x(P,`
            Price change in 24 hours`),P.forEach(o),pe.forEach(o),z=y(K),S=v(K,"TH",{});var R=_(S);C=v(R,"DIV",{class:!0});var A=_(C);j=x(A,"Min/Max price"),A.forEach(o),c=y(R),d=v(R,"DIV",{class:!0});var D=_(d);Q=x(D,"In 24 hours"),D.forEach(o),R.forEach(o),X=y(K),J=v(K,"TH",{});var ve=_(J);se=v(ve,"DIV",{class:!0});var ge=_(se);le=x(ge,"Current price meme"),ge.forEach(o),me=y(ve),Z=v(ve,"DIV",{class:!0});var we=_(Z);ue=x(we,"Generated hourly"),we.forEach(o),ve.forEach(o),K.forEach(o),Ae.forEach(o),_e=y($),G=v($,"TBODY",{});var Ee=_(G);F.l(Ee),Ee.forEach(o),$.forEach(o),re.forEach(o),this.h()},h(){i(g,"class","normal-case font-black tracking-wider"),i(N,"d","M17.92,6.62a1,1,0,0,0-.54-.54A1,1,0,0,0,17,6H7A1,1,0,0,0,7,8h7.59l-8.3,8.29a1,1,0,0,0,0,1.42,1,1,0,0,0,1.42,0L16,9.41V17a1,1,0,0,0,2,0V7A1,1,0,0,0,17.92,6.62Z"),i(I,"class","w-3 h-3 fill-current"),i(I,"xmlns","http://www.w3.org/2000/svg"),i(I,"viewBox","0 0 24 24"),i(L,"class","flex items-center gap-1 text-xs normal-case opacity-70"),i(C,"class","normal-case font-black tracking-wider"),i(d,"class","text-xs normal-case opacity-70"),i(se,"class","normal-case font-black tracking-wider"),i(Z,"class","text-xs normal-case opacity-70"),i(t,"class","table table-zebra w-full"),i(l,"class","overflow-x-auto")},m(E,W){Ve(e,document.head,null),b(E,r,W),b(E,l,W),u(l,t),u(t,s),u(s,n),u(n,h),B&&B.m(h,null),u(n,m),u(n,f),u(f,g),u(g,w),u(f,T),u(f,L),u(L,I),u(I,N),u(L,M),u(n,z),u(n,S),u(S,C),u(C,j),u(S,c),u(S,d),u(d,Q),u(n,X),u(n,J),u(J,se),u(se,le),u(J,me),u(J,Z),u(Z,ue),u(t,_e),u(t,G),ee[H].m(G,null),he=!0},p(E,[W]){ae===(ae=ie(E))&&B?B.p(E,W):(B&&B.d(1),B=ae&&ae(E),B&&(B.c(),B.m(h,null)));let re=H;H=ne(E),H===re?ee[H].p(E,W):(ze(),te(ee[re],1,1,()=>{ee[re]=null}),Le(),F=ee[H],F?F.p(E,W):(F=ee[H]=de[H](E),F.c()),q(F,1),F.m(G,null))},i(E){he||(q(e.$$.fragment,E),q(F),he=!0)},o(E){te(e.$$.fragment,E),te(F),he=!1},d(E){xe(e),E&&o(r),E&&o(l),B&&B.d(),ee[H].d()}}}function Lt(a,e,r){let l;l="https://api.cryptomeme.wtf/v1";let t,s,n="idle";return qe(()=>{function h(){r(2,n="loading"),r(0,t=JSON.parse(localStorage.getItem("coins"))),r(1,s=localStorage.getItem("lastUpdate")),fetch(l+"/markets").then(f=>f.json()).then(f=>{r(2,n="success"),r(0,t=f),localStorage.setItem("coins",JSON.stringify(t)),localStorage.setItem("lastUpdate",new Date().toISOString())}).catch(f=>{r(2,n="error")})}const m=setInterval(h,10*1e3);return h(),()=>clearInterval(m)}),[t,s,n]}class Ot extends Te{constructor(e){super();Ie(this,e,Lt,Pt,Se,{})}}export{Ot as default};
