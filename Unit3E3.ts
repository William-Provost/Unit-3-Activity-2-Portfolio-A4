/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-09
 * @fileoverview This program prompts the user to input four lines of text.
 * It then displays the lines in reverse order (line 4, line 3, line 2, line 1).
 */

// variables
let line1: string;
let line2: string;
let line3: string;
let line4: string;

// input
line1 = prompt("Enter line 1:") || "";
line2 = prompt("Enter line 2:") || "";
line3 = prompt("Enter line 3:") || "";
line4 = prompt("Enter line 4:") || "";

// output
console.log("\nThe lines in reverse order are:");
console.log(line4);
console.log(line3);
console.log(line2);
console.log(line1);
console.log("Done.");
