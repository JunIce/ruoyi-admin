import { createI18n } from "vue-i18n";
import zh_cn from "./locale/zh_cn.json";
import en from "./locale/en.json";

const i18n = createI18n({
  // something vue-i18n options here ...
  locale: "zh_cn",
  fallbackLocale: "zh_cn",
  messages: {
    zh_cn: zh_cn,
    en: en,
  },
});

export default i18n;
