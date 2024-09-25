export interface Matcher {
  //returns true if the char can be consumed by the matcher
  matcher(char: string): boolean;
  //need to knwo when it is an epsilon transition to avoid consuming input
  isEpsilon(): null;

  label(): string;
}
