import type { Matcher } from "./matcher";

export class EpsilonMatcher implements Matcher {
  matches(char: string): boolean {
    return true;
  }

  isEpsilon(): boolean {
    return true;
  }

  label(): string {
    return "ε";
  }
}
