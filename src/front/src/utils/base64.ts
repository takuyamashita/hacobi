export const base64URLSafeToUint8Array = (base64URLSafe: string) => {
  if (base64URLSafe === undefined || base64URLSafe === null) {
    return new Uint8Array();
  }

  // padding
  const pad = (s: string) => {
    while (s.length % 4 !== 0) {
      s += "=";
    }
    return s;
  };

  // base64URLSafe to base64
  const base64 = pad(base64URLSafe).replace(/\-/g, "+").replace(/_/g, "/");

  // base64 to Uint8Array
  const raw = window.atob(base64);
  const rawLength = raw.length;
  const array = new Uint8Array(new ArrayBuffer(rawLength));
  for (let i = 0; i < rawLength; i++) {
    array[i] = raw.charCodeAt(i);
  }
  return array;
};

/**
 * paddingありのbase64URLSafeに変換する
 *
 * @param arrayBuffer
 * @returns paddingありのbase64URLSafe
 */
export const arrayBufferToBase64URLSafe = (arrayBuffer: ArrayBuffer): string =>
  window
    .btoa(String.fromCharCode(...new Uint8Array(arrayBuffer)))
    .replace(/\+/g, "-")
    .replace(/\//g, "_")
    .replace(/=/g, "");
