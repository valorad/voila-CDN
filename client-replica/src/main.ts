import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import "./registerServiceWorker";

import {
  MdButton,
  MdContent,
  MdTabs,
  MdApp,
  MdDrawer,
  MdToolbar,
  MdList,
  MdIcon,
  MdTable,
  MdField,
  MdCard,
  MdCheckbox,
  MdDialog,
  MdSteppers,
  MdMenu,
} from 'vue-material/dist/components';

Vue.config.productionTip = false;

// vue material
Vue.use(MdButton);
Vue.use(MdTabs);
Vue.use(MdApp);
Vue.use(MdDrawer);
Vue.use(MdToolbar);
Vue.use(MdList);
Vue.use(MdIcon);
Vue.use(MdTable);
Vue.use(MdField);
Vue.use(MdCard);
Vue.use(MdCheckbox);
Vue.use(MdDialog);
Vue.use(MdSteppers);
Vue.use(MdContent);
Vue.use(MdMenu);


new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
