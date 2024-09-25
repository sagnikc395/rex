import type { Matcher } from "./matcher";

export class CharacterMatcher implements Matcher {
  c: string;
  constructor(c: string) {
    this.c = c;
  }

  matches(char: string) {
    return this.c === char;
  }

  isEpsilon(): boolean {
    return false;
  }

  label(): string {
    return this.c;
  }
}
