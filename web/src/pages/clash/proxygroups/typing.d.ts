export interface Node {
  nodeType: string;

  server?: string;
  remarks?: string;
  port?: number;
  encryption?: string;
  password?: string;
  plugin?: string;
  plugin_options?: Object;

  //   server: string;
  //   remarks: string;
  //   port: number;
  //   encryption: string;
  //   password: string;
  protocol?: string;
  protocol_param?: string;
  obfs?: string;
  obfs_param?: string;
  group?: string;

  host?: string;
  path?: string;
  tls?: string;
  add?: string;
  //   port: number;
  aid?: number;
  net?: string;
  type?: string;
  v?: string;
  ps?: string;
  id?: string;
  class?: number;
}

export interface Group {
  name: string;
  url: string;
  nodes: Node[];
  lastUpdate: string;
  updating: boolean;
}

export interface ProxySelector {
  groupName: string;
  include: string;
  exclude: string;
  selected: boolean;
}

export interface ClashProxyGroupSelector {
  name: string;
  type: string;
  url: string;
  interval?: number;
  proxyGroups: string[];
  proxySelectors: ProxySelector[];
  proxies: Node[];
}
