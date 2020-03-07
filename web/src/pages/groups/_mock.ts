import { Request, Response } from 'express';
import { Group, Node } from './typing';
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

function generateGroup() {
  const result: {
    name: string;
    url: string;
    nodes: Node[];
    lastUpdate: string;
  } = {
    name: Random.string(),
    url: Random.integer(0, 1) ? '' : Random.url('https'),
    nodes: [],
    lastUpdate: Random.datetime(),
  };
  for (let i = 0; i < 50; i++) {
    let rd = Random.integer(0, 2);
    if (rd == 0) {
      result.nodes.push(generateSS());
    } else if (rd == 1) {
      result.nodes.push(generateSSR());
    } else if (rd == 2) {
      result.nodes.push(generateVmess());
    }
  }
  return result;
}

function generateGroups(req: Request, res: Response) {
  const result = [];
  for (let i = 0; i < 10; i++) {
    result.push(generateGroup());
  }
  return res.json({
    status: 'success',
    data: result,
  });
}

function updateGroup(req: Request, res: Response) {
  const { name } = req.body;
  const result = generateGroup();
  result.name = name;
  result.lastUpdate = Random.now('second');
  setTimeout(() => {
    return res.json({
      status: 'success',
      data: result,
    });
  }, 1200);
}

function addOrEditNode(req: Request, res: Response) {
  let groupName = req.query.group;
  let group = generateGroup();
  group.name = groupName;
  if (Random.integer(0, 1)) {
    return res.json({
      status: 'success',
      data: group,
    });
  }
  return res.json({
    status: 'fail',
    msg: 'error',
  });
}

function deleteGroup(req: Request, res: Response) {
  if (Random.integer(0, 1)) {
    generateGroups(req, res);
  } else {
    return res.json({
      status: 'fail',
      msg: 'error',
    });
  }
}

function deleteNode(req: Request, res: Response) {
  if (Random.integer(0, 1)) {
    let groupName = req.query.group;
    let group = generateGroup();
    group.name = groupName;
    return res.json({
      status: 'success',
      data: group,
    });
  }
  return res.json({
    status: 'fail',
    msg: 'error',
  });
}

function addGroup(req: Request, res: Response) {
  setTimeout(() => {
    if (Random.integer(0, 1)) {
      generateGroups(req, res);
    } else {
      return res.json({
        status: 'fail',
        msg: 'error',
      });
    }
  }, 1200);
}

export default {
  'GET /api/groups': generateGroups,
  'POST /api/group': (req: Request, res: Response) => {
    let action = req.query.action;
    action === 'add' && addGroup(req, res);
    action === 'edit' && generateGroups(req, res);
    action === 'update' && updateGroup(req, res);
    action === 'delete' && deleteGroup(req, res);
    action === 'updateall' && generateGroups(req, res);
  },
  'POST /api/node': (req: Request, res: Response) => {
    let action = req.query.action;
    action === 'add' && addOrEditNode(req, res);
    action === 'edit' && addOrEditNode(req, res);
    action === 'delete' && deleteNode(req, res);
  },
};
