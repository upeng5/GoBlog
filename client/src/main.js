import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Vuesax from "vuesax";
import "vuesax/dist/vuesax.css"; //Vuesax styles
import Unicon from "vue-unicons";
import { uniLayerGroupMonochrome, uniCarWash } from "vue-unicons/src/icons";

Vue.config.productionTip = false;

Unicon.add([uniLayerGroupMonochrome, uniCarWash]);
Vue.use(Unicon);
Vue.use(Vuesax);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
