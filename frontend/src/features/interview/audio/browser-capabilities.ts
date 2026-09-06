export function getBrowserCapabilities() {
  if (typeof window === "undefined")
    return {
      microphone: false,
      camera: false,
      mediaRecorder: false,
      webAudio: false,
      webGL: false,
      online: false,
    };
  const canvas = document.createElement("canvas");
  const gl = canvas.getContext("webgl2") ?? canvas.getContext("webgl");
  const webGL = gl !== null;
  gl?.getExtension("WEBGL_lose_context")?.loseContext();
  return {
    microphone: Boolean(navigator.mediaDevices?.getUserMedia),
    camera: Boolean(navigator.mediaDevices?.getUserMedia),
    mediaRecorder: typeof MediaRecorder !== "undefined",
    webAudio: typeof window.AudioContext !== "undefined",
    webGL,
    online: navigator.onLine,
  };
}
