export interface RuleGroup {
  name: string;
  proxyGroup: string;
  url: string;
  rules: string[];
  customRules: string[];
  lastUpdate: string;
}
