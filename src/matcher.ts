export interface Matcher {
  //returns true if the char can be consumed by the matcher
  matches(char: string): boolean;
  //need to knwo when it is an epsilon transition to avoid consuming input
  isEpsilon(): boolean;

  label(): string;
}
