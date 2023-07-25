function replaceCharacters() {
    const inputText = document.getElementById("inputText").value;
    const result = replaceCharactersJS(inputText);
    document.getElementById("output").innerText = result;
}

function replaceCharactersJS(input) {
    const charMap = {
        'a': '𝕒', 'b': '𝕓', 'c': '𝕔', 'd': '𝕕', 'e': '𝕖', 'f': '𝕗', 'g': '𝕘', 'h': '𝕙', 'i': '𝕚', 'j': '𝕛', 'k': '𝕜',
        'l': '𝕝', 'm': '𝕞', 'n': '𝕟', 'o': '𝕠', 'p': '𝕡', 'q': '𝕢', 'r': '𝕣', 's': '𝕤', 't': '𝕥', 'u': '𝕦', 'v': '𝕧',
        'w': '𝕨', 'x': '𝕩', 'y': '𝕪', 'z': '𝕫', 'A': '𝔸', 'B': '𝔹', 'C': 'ℂ', 'D': '𝔻', 'E': '𝔼', 'F': '𝔽', 'G': '𝔾',
        'H': 'ℍ', 'I': '𝕀', 'J': '𝕁', 'K': '𝕂', 'L': '𝕃', 'M': '𝕄', 'N': 'ℕ', 'O': '𝕆', 'P': 'ℙ', 'Q': 'ℚ', 'R': 'ℝ',
        'S': '𝕊', 'T': '𝕋', 'U': '𝕌', 'V': '𝕍', 'W': '𝕎', 'X': '𝕏', 'Y': '𝕐', 'Z': 'ℤ', '0': '𝟘', '1': '𝟙', '2': '𝟚',
        '3': '𝟛', '4': '𝟜', '5': '𝟝', '6': '𝟞', '7': '𝟟', '8': '𝟠', '9': '𝟡',
    };
    let result = "";
    for (const char of input) {
        if (charMap[char] !== undefined) {
            result += charMap[char];
        } else {
            result += char;
        }
    }
    return result;
}

window.onload = function() {
    replaceCharacters();
};

// Copy to Clipboard
function copyToClipboard() {
    const outputText = document.getElementById("output").innerText;
    const tempInput = document.createElement("textarea");
    tempInput.style = "position: absolute; left: -1000px; top: -1000px";
    tempInput.value = outputText;
    document.body.appendChild(tempInput);
    tempInput.select();
    document.execCommand("copy");
    document.body.removeChild(tempInput);
}

function shareTwitter() {
    const result = document.getElementById('output').textContent;
    const url = 'https://x.shiyui.dev/';
    const textWithUrl = `${result}\n${url}`;
    const encodedResult = encodeURIComponent(textWithUrl);
    const twitterShareURL = `https://twitter.com/intent/tweet?text=${encodedResult}`;
    window.open(twitterShareURL, '_blank');
}
