(function(t){function e(e){for(var o,i,r=e[0],l=e[1],c=e[2],d=0,p=[];d<r.length;d++)i=r[d],Object.prototype.hasOwnProperty.call(n,i)&&n[i]&&p.push(n[i][0]),n[i]=0;for(o in l)Object.prototype.hasOwnProperty.call(l,o)&&(t[o]=l[o]);u&&u(e);while(p.length)p.shift()();return a.push.apply(a,c||[]),s()}function s(){for(var t,e=0;e<a.length;e++){for(var s=a[e],o=!0,r=1;r<s.length;r++){var l=s[r];0!==n[l]&&(o=!1)}o&&(a.splice(e--,1),t=i(i.s=s[0]))}return t}var o={},n={app:0},a=[];function i(e){if(o[e])return o[e].exports;var s=o[e]={i:e,l:!1,exports:{}};return t[e].call(s.exports,s,s.exports,i),s.l=!0,s.exports}i.m=t,i.c=o,i.d=function(t,e,s){i.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:s})},i.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},i.t=function(t,e){if(1&e&&(t=i(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var s=Object.create(null);if(i.r(s),Object.defineProperty(s,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)i.d(s,o,function(e){return t[e]}.bind(null,o));return s},i.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return i.d(e,"a",e),e},i.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},i.p="/";var r=window["webpackJsonp"]=window["webpackJsonp"]||[],l=r.push.bind(r);r.push=e,r=r.slice();for(var c=0;c<r.length;c++)e(r[c]);var u=l;a.push([0,"chunk-vendors"]),s()})({0:function(t,e,s){t.exports=s("56d7")},"02b1":function(t,e,s){"use strict";var o=s("7f46"),n=s.n(o);n.a},"034f":function(t,e,s){"use strict";var o=s("85ec"),n=s.n(o);n.a},"1a12":function(t,e,s){},"1ec6":function(t,e,s){"use strict";var o=s("fa57"),n=s.n(o);n.a},"2cfd":function(t,e,s){"use strict";var o=s("e314"),n=s.n(o);n.a},"56d7":function(t,e,s){"use strict";s.r(e);s("e260"),s("e6cf"),s("cca6"),s("a79d");var o=s("2b0e"),n=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{attrs:{id:"app"}},[s("Navbar"),s("router-view")],1)},a=[],i=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"navbar"},[s("div",{staticClass:"container"},[s("router-link",{staticClass:"navbar-brand",attrs:{to:"/"}},[s("i",{staticClass:"material-icons book"},[t._v("book")]),t._v(" GoBlog ")]),s("div",{staticClass:"nav-menu"},[t.status?s("ul",[s("li",[s("router-link",{attrs:{to:"/"}},[s("vs-button",{attrs:{shadow:""},scopedSlots:t._u([{key:"animate",fn:function(){return[s("i",{staticClass:"material-icons"},[t._v("home")])]},proxy:!0}],null,!1,1613395375)},[t._v(" Home ")])],1)],1),s("li",[s("vs-button",{attrs:{danger:"",shadow:""},on:{click:t.logout}},[t._v("Logout")])],1)]):s("ul",[s("li",[s("router-link",{attrs:{to:"/login"}},[s("vs-button",{attrs:{shadow:""}},[t._v("Login")])],1)],1),s("li",[s("router-link",{attrs:{to:"/register"}},[s("vs-button",{attrs:{shadow:""}},[t._v("Register")])],1)],1)])])],1)])},r=[],l={name:"Navbar",computed:{status:function(){return this.$store.state.status}},methods:{logout:function(){this.$store.dispatch("logout"),this.$store.state.status=!1,this.$router.push("/login")}}},c=l,u=(s("5dfc"),s("2877")),d=Object(u["a"])(c,i,r,!1,null,null,null),p=d.exports,f={name:"App",components:{Navbar:p}},m=f,h=(s("034f"),Object(u["a"])(m,n,a,!1,null,null,null)),v=h.exports,g=s("8c4f"),b=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"home"},[s("FloatingAdd"),t._l(t.posts,(function(t,e){return s("div",{key:e},[s("Post",{attrs:{post:t}})],1)}))],2)},w=[],_=(s("b0c0"),s("d3b7"),s("2f62")),k=s("04e1"),y=s.n(k),C=s("bc3a"),x=s.n(C);o["default"].use(_["a"]);var P=new _["a"].Store({state:{status:!1,user:{token:null,user_name:null}},getters:{getUser:function(t){return t.user},getStatus:function(t){return t.isAuthorized}},mutations:{postLogin:function(t,e){var s=e.email,o=e.password;fetch("http://localhost:3000/api/v1/users",{method:"POST",mode:"cors",headers:{"Content-Type":"application/json"},body:JSON.stringify({email:s,password:o})}).then((function(e){e.json().then((function(e){var s=y()(e);console.log(s),t.user.token=e,t.user.user_name=s.name,t.status=!0,ut.push("/")})).catch((function(t){console.log(t)}))}))},logout:function(t){t.isAuthorized=!1,t.user.token=null,t.user.user_name=null},postRegister:function(t,e){var s=e.username,o=e.email,n=e.password;x.a.post("http://localhost:3000/api/v1/users/register",{username:s,email:o,password:n}).then((function(e){var s=e.data,o=y()(s.token);console.log(o),t.user.token=s.token,t.user.user_name=o.name,t.status=!0,ut.push("/")})).catch((function(t){console.log(t)}))}},actions:{postLogin:function(t,e){var s=e.email,o=e.password;t.commit("postLogin",{email:s,password:o})},logout:function(t){t.commit("logout")},postRegister:function(t,e){var s=e.username,o=e.email,n=e.password;t.commit("postRegister",{username:s,email:o,password:n})}},modules:{}}),S=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"floating-add"},[s("vs-button",{staticClass:"f-button",attrs:{circle:"",icon:"",floating:"",color:"#ffba00"},on:{click:t.loadAddPostPage}},[s("i",{staticClass:"material-icons"},[t._v("add")])])],1)},$=[],O={name:"FloatingAdd",methods:{loadAddPostPage:function(){this.$router.push("/add-post")}}},j=O,E=(s("8ee7"),Object(u["a"])(j,S,$,!1,null,null,null)),A=E.exports,M=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"post-component"},[s("div",{staticClass:"post"},[s("vs-row",[s("vs-col",{staticClass:"avatar-col",attrs:{w:"2"}},[s("vs-avatar",{attrs:{warn:"",circle:"",size:"60"},scopedSlots:t._u([{key:"text",fn:function(){return[t._v(t._s(t.post.username))]},proxy:!0}])})],1),s("vs-col",{attrs:{w:"10"}},[s("h3",{staticClass:"post-title"},[t._v(t._s(t.post.title))]),s("p",[t._v(t._s(t.post.content))])]),t.user.user_name===t.post.username?s("vs-button",{staticClass:"f-button",attrs:{circle:"",icon:"",floating:"",color:"#ffba00"},on:{click:t.deletePost}},[s("i",{staticClass:"material-icons"},[t._v("delete")])]):t._e()],1)],1)])},N=[],R={name:"Post",props:["post"],computed:{user:function(){return this.$store.state.user}},methods:{deletePost:function(){var t=this.post._id;console.log(t),x.a.delete("http://localhost:3000/api/v1/posts/".concat(t),{headers:{"x-auth-token":this.user.token}}).then((function(t){console.log("Delete Request:",t)})).catch((function(t){console.log(t)}))}}},T=R,L=(s("02b1"),Object(u["a"])(T,M,N,!1,null,null,null)),z=L.exports,F={name:"Home",components:{FloatingAdd:A,Post:z},data:function(){return{posts:[],loading:!0}},store:P,methods:{getPosts:function(){var t=this;x.a.get("http://localhost:3000/api/v1/posts",{headers:{"x-auth-token":this.userToken}}).then((function(e){200===e.status&&(console.log(e.data),t.posts=e.data),console.log(e)})).catch((function(t){console.log(t)}))}},computed:{userToken:function(){return this.$store.state.user.token}},created:function(){},mounted:function(){this.getPosts(),this.loading=!1}},H=F,J=(s("cccb"),Object(u["a"])(H,b,w,!1,null,null,null)),U=J.exports,q=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"login"},[s("vs-row",[s("vs-col",{attrs:{offset:"5",w:"6"}},[s("vs-input",{staticClass:"email-field",attrs:{warn:"",placeholder:"Email"},scopedSlots:t._u([{key:"icon",fn:function(){return[s("i",{staticClass:"material-icons"},[t._v("email")])]},proxy:!0}]),model:{value:t.email,callback:function(e){t.email=e},expression:"email"}}),s("vs-input",{staticClass:"password-field",attrs:{warn:"",type:"password",placeholder:"Password"},scopedSlots:t._u([{key:"icon",fn:function(){return[s("i",{staticClass:"material-icons"},[t._v("lock")])]},proxy:!0}]),model:{value:t.password,callback:function(e){t.password=e},expression:"password"}}),s("vs-button",{staticClass:"submit-button",attrs:{warn:""},on:{click:t.loginAction},scopedSlots:t._u([{key:"animate",fn:function(){return[s("i",{staticClass:"material-icons"},[t._v("login")])]},proxy:!0}])},[t._v(" Submit ")])],1)],1)],1)},B=[],D=s("5b7e"),G=s.n(D),I={name:"Login",store:P,components:{Spinner:G.a},data:function(){return{email:"",password:"",errorMessage:"",loading:!1}},computed:{isAuth:function(){return this.$store.state.status}},methods:{loginAction:function(){if(this.loading=!0,0===this.email.length||0===this.password.length)return this.errorMessage="Please enter valid credentials",this.openNotification(),void(this.loading=!1);var t={email:this.email,password:this.password};this.$store.dispatch("postLogin",t),this.loading=!1},openNotification:function(){this.$vs.notification({progress:"auto",border:"warm",position:"top-left",title:"Error",text:this.errorMessage})}}},K=I,Q=(s("d6db"),Object(u["a"])(K,q,B,!1,null,null,null)),V=Q.exports,W=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"register center"},[s("vs-row",[s("vs-col",{attrs:{offset:"5",w:"6"}},[s("vs-input",{staticClass:"username-field",attrs:{warn:"",placeholder:"Username"},scopedSlots:t._u([{key:"icon",fn:function(){return[s("i",{staticClass:"material-icons"},[t._v("person")])]},proxy:!0}]),model:{value:t.username,callback:function(e){t.username=e},expression:"username"}}),s("vs-input",{staticClass:"email-field",attrs:{warn:"",placeholder:"Email"},scopedSlots:t._u([{key:"icon",fn:function(){return[s("i",{staticClass:"material-icons"},[t._v("email")])]},proxy:!0}]),model:{value:t.email,callback:function(e){t.email=e},expression:"email"}}),s("vs-input",{staticClass:"password-field",attrs:{warn:"",type:"password",placeholder:"Password"},scopedSlots:t._u([{key:"icon",fn:function(){return[s("i",{staticClass:"material-icons"},[t._v("lock")])]},proxy:!0}]),model:{value:t.password,callback:function(e){t.password=e},expression:"password"}}),s("vs-button",{staticClass:"submit-button",attrs:{warn:""},on:{click:t.postRegister},scopedSlots:t._u([{key:"animate",fn:function(){return[s("i",{staticClass:"material-icons"},[t._v("login")])]},proxy:!0}])},[t._v(" Submit ")])],1)],1)],1)},X=[],Y={name:"Register",components:{Spinner:G.a},data:function(){return{username:"",email:"",password:""}},methods:{postRegister:function(){return 0===this.username.length?(this.errorMessage="Please enter a valid username",this.openNotification(),void(this.loading=!1)):0===this.email.length?(this.errorMessage="Please enter a valid email",this.openNotification(),void(this.loading=!1)):0===this.password.length?(this.errorMessage="Please enter a valid password",this.openNotification(),void(this.loading=!1)):void this.$store.dispatch("postRegister",{username:this.username,email:this.email,password:this.password})},openNotification:function(){this.$vs.notification({progress:"auto",border:"warm",position:"top-left",title:"Error",text:this.errorMessage})}}},Z=Y,tt=(s("2cfd"),Object(u["a"])(Z,W,X,!1,null,null,null)),et=tt.exports,st=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"add-post"},[s("h1",[t._v("Add Post")]),s("vs-row",[s("vs-col",{staticClass:"add-form"},[s("div",[s("vs-input",{staticClass:"title-field",attrs:{warn:"",border:"",placeholder:"Title"},scopedSlots:t._u([{key:"icon",fn:function(){return[s("i",{staticClass:"material-icons"},[t._v("title")])]},proxy:!0}]),model:{value:t.title,callback:function(e){t.title=e},expression:"title"}}),s("textarea",{directives:[{name:"model",rawName:"v-model",value:t.content,expression:"content"}],staticClass:"content-field",attrs:{cols:"50",rows:"10"},domProps:{value:t.content},on:{input:function(e){e.target.composing||(t.content=e.target.value)}}}),s("vs-button",{staticClass:"submit-button",attrs:{warn:""},on:{click:t.createPost},scopedSlots:t._u([{key:"animate",fn:function(){return[s("i",{staticClass:"material-icons"},[t._v("add")])]},proxy:!0}])},[t._v(" Submit ")])],1)])],1)],1)},ot=[],nt=(s("498a"),{name:"AddPost",data:function(){return{title:"",content:""}},computed:{userToken:function(){return this.$store.state.user.token}},methods:{createPost:function(){var t=this;x.a.post("http://localhost:3000/api/v1/posts",{title:this.title,content:this.content.trim()},{headers:{"x-auth-token":this.userToken}}).then((function(e){console.log(e),t.$router.push("/")})).catch((function(t){console.log(t)}))}}}),at=nt,it=(s("1ec6"),Object(u["a"])(at,st,ot,!1,null,null,null)),rt=it.exports;o["default"].use(g["a"]);var lt=[{path:"/",name:"Home",component:U,beforeEnter:function(t,e,s){!1===P.state.status?s("/login"):s()}},{path:"/login",name:"Login",component:V,beforeEnter:function(t,e,s){!0===P.state.status?s("/"):s()}},{path:"/register",name:"Register",component:et,beforeEnter:function(t,e,s){!0===P.state.status?s("/"):s()}},{path:"/add-post",name:"Add Post",component:rt,beforeEnter:function(t,e,s){!1===P.state.status?s("/login"):s()}}],ct=new g["a"]({routes:lt}),ut=ct,dt=s("574d"),pt=s.n(dt),ft=(s("04f2"),s("a2e0")),mt=s.n(ft),ht=s("cf18");o["default"].config.productionTip=!1,mt.a.add([ht["b"],ht["a"]]),o["default"].use(mt.a),o["default"].use(pt.a),new o["default"]({router:ut,store:P,render:function(t){return t(v)}}).$mount("#app")},"5ced":function(t,e,s){},"5dfc":function(t,e,s){"use strict";var o=s("1a12"),n=s.n(o);n.a},"7f46":function(t,e,s){},"85ec":function(t,e,s){},"8ee7":function(t,e,s){"use strict";var o=s("fd72"),n=s.n(o);n.a},cccb:function(t,e,s){"use strict";var o=s("5ced"),n=s.n(o);n.a},d6db:function(t,e,s){"use strict";var o=s("e67a"),n=s.n(o);n.a},e314:function(t,e,s){},e67a:function(t,e,s){},fa57:function(t,e,s){},fd72:function(t,e,s){}});
//# sourceMappingURL=app.1bb1f939.js.map