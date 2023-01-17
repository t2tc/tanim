
const Content = {
    "zh-Hans": {},
    "en": {}
}

type Language = "zh-Hans" | "ja" | "en";
const language: Language = navigator.language === "zh-CN" ? "zh-Hans" :
    navigator.language === "zh-SG" ? "zh-Hans" :
        navigator.language === "ja" ? "ja" :
            "en";

(() => {
    
})();