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

  username?: string;
  skipCertVerify?: boolean;
}

export interface Group {
  name: string;
  url: string;
  nodes: Node[];
  lastUpdate: string;
  updating: boolean;
}
