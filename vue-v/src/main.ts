import { createApp } from "vue";
import App from "./App.vue";
import { createRouter, createWebHistory } from "vue-router";
import Manage from "./components/Manage.vue";
import Notes from "./components/Notes.vue";
import "./index.css";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", redirect: "/manage" },
    { path: "/manage", component: Manage },
    { path: "/notes", component: Notes },
  ],
});

const app = createApp(App);

app.use(router);
app.mount("#app");
