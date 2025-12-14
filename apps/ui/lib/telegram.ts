export function getTelegram() {
    if (typeof window === "undefined") return null;
    return (window as any).Telegram?.WebApp || null;
}


export function getInitData(): string {
    return getTelegram()?.initData || "";
}