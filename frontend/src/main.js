import { createApp, provide, h } from "vue";
import { DefaultApolloClient } from "@vue/apollo-composable";
import { apolloProvider } from "./apollo";
import App from "./App.vue";
import "./index.css";

const app = createApp({
  render: () => h(App),
});

app.use(apolloProvider);

app.mount("#app");
