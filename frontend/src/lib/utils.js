export const handleCopyToClipboard = (data) => {
  navigator.clipboard
    .writeText(data)
    .then(() => {
      alert("ShortURL copied to clipboard!");
    })
    .catch((err) => {
      console.error("Failed to copy text: ", err);
    });
};
