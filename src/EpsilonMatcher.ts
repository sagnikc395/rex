import type { Matcher } from "./Matcher";

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
