import { ref, watch } from "vue";

export function useSysmode() {
  const darkMode = ref(false);

  const handleDarkModeChange = () => {
    console.log("handleDarkModeChange");

    darkMode.value = !darkMode.value;
  };

  watch(darkMode, (newVal) => {
    if (newVal) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  });

  return {
    darkMode,
    handleDarkModeChange,
  };
}
