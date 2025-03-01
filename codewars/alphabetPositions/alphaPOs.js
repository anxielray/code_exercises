function alphabetPosition(text) {
  let result = "";
  const arr = text.split(" ");

  for (let i = 0; i < arr.length; i++) {
    const word = arr[i];

    for (let x = 0; x < word.length; x++) {
      const c = word[x];

      if (c >= "A" && c <= "Z") {
        const value = c.charCodeAt(0) - 64; // Convert uppercase letter to position
        result += i !== arr.length - 1 ? value + " " : value;
      } else if (c >= "a" && c <= "z") {
        const value = c.charCodeAt(0) - 96; // Convert lowercase letter to position
        if (i !== arr.length - 1) {
          result += value + " ";
        } else if (i === arr.length - 1 && x === c.length - 1) {
          result += value;
        }
      }
    }
  }
  return result;
}

console.log(alphabetPosition("The sunset sets at twelve o' clock."))