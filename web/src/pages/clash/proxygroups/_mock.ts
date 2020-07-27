import { Request, Response, json } from 'express';
import { Node, ClashProxyGroupSelector, ProxySelector } from './typing';
import Mock from 'mockjs';

const Random = Mock.Random;

function generateSS() {
  return {
    nodeType: 'ss',
    server: `${Random.domain()}`,
    remarks: `${Random.region()} ${Random.county()} ${Random.integer(1, 10)}`,
    port: Random.integer(1024, 65536),
    encryption: 'aes-256-cfb',
    password: `${Random.sentence(1)}`,
  };
}

function generateSSR() {
  return {
    nodeType: 'ssr',
    server: `${Random.domain()}`,
    remarks: `${Random.region()} ${Random.county()} ${Random.integer(1, 10)}`,
    port: Random.integer(1024, 65536),
    encryption: 'aes-256-cfb',
    password: `${Random.sentence(1)}`,
    protocol: 'origin',
    protocol_param: `${Random.sentence(1)}`,
    obfs: 'plain',
    obfs_param: `${Random.sentence(1)}`,
    group: `${Random.sentence(1)}`,
  };
}

function generateVmess() {
  return {
    nodeType: 'vmess',
    host: `${Random.domain()}`,
    path: `/${Random.string()}`,
    tls: `tls`,
    add: `${Random.domain()}`,
    port: Random.integer(1025, 65535),
    aid: 2,
    net: `${Random.sentence(1)}`,
    type: `auto`,
    v: `${Random.sentence(1)}`,
    ps: `${Random.region()} ${Random.county()} ${Random.integer(1, 10)}`,
    id: `${Random.guid()}`,
    class: 1,
  };
}

function generateType() {
  let type = Mock.mock({
    'type|1': ['url-test', 'fallback', 'load-balance', 'select'],
  });
  return type.type;
}

function generateProxySelector() {
  let result: ProxySelector = {
    groupName: Random.ctitle(),
    include: '',
    exclude: '',
    selected: Random.integer(0, 1) ? true : false,
  };
  let count = Random.integer(5, 20);
  for (let i = 0; i < count; i++) {
    i === count - 1
      ? (result.include += Random.county())
      : (result.include += Random.county() + '|');
    i === count - 1
      ? (result.exclude += Random.county())
      : (result.exclude += Random.county() + '|');
  }
  return result;
}

function generateSelector() {
  let type = generateType();
  let selector: ClashProxyGroupSelector = {
    name: Random.ctitle(),
    type: type,
    url: type === 'select' ? '' : Random.url('https'),
    interval: type === 'select' ? 0 : Random.integer(200, 800),
    proxyGroups: [],
    proxySelectors: [],
    proxies: [],
  };
  let count = Random.integer(5, 20);
  for (let i = 0; i < count; ++i) {
    selector.proxyGroups.push(Random.ctitle());
  }
  count = Random.integer(5, 20);
  for (let i = 0; i < count; ++i) {
    selector.proxySelectors.push(generateProxySelector());
  }
  count = Random.integer(5, 20);
  for (let i = 0; i < count; ++i) {
    let node: Node;
    let rd = Random.integer(0, 2);
    switch (rd) {
      case 0:
        node = generateSS();
        break;
      case 1:
        node = generateSSR();
        break;
      case 2:
        node = generateVmess();
        break;
      default:
        node = generateSS();
        break;
    }
    selector.proxies.push(node);
  }
  return selector;
}

function generateSelectors(req: Request, res: Response) {
  let result: ClashProxyGroupSelector[] = [];
  let count = Random.integer(5, 10);
  for (let i = 0; i < count; ++i) {
    result.push(generateSelector());
  }
  return res.json({
    status: 'success',
    data: result,
  });
}

function addSelector(req: Request, res: Response) {
  if (Random.integer(0, 1)) {
    return res.json({
      status: 'fail',
      msg: 'error',
    });
  }
  generateSelectors(req, res);
}

function updateSelector(req: Request, res: Response) {
  let { name } = req.body;
  let result = generateSelector();
  result.name = name;
  return res.json({
    status: 'success',
    data: result,
  });
}

function editSelector(req: Request, res: Response) {
  if (Random.integer(0, 1)) {
    return res.json({
      status: 'fail',
      msg: 'error',
    });
  }
  generateSelectors(req, res);
}

export default {
  'GET /api/selectors': generateSelectors,
  'POST /api/selector': (req: Request, res: Response) => {
    let action = req.query.action;
    action === 'add' && addSelector(req, res);
    action === 'edit' && editSelector(req, res);
    action === 'update' && updateSelector(req, res);
    action === 'delete' && editSelector(req, res);
    action === 'updateall' && generateSelectors(req, res);
  },
};
