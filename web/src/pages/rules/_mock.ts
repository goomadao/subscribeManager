import { Request, Response } from 'express';
import { RuleGroup } from './typing';
import Mock from 'mockjs';

const Random = Mock.Random;

function generateRule(name?: string, url?: string) {
  let haveURL = Random.integer(0, 1);
  let result: RuleGroup = {
    name: name || Random.ctitle(),
    proxyGroup: Random.ctitle(),
    url: url || (haveURL ? Random.url('https') : ''),
    rules: [],
    customRules: [],
    lastUpdate: Random.datetime(),
  };
  let count = Random.integer(10, 50);
  for (let i = 0; i < count; i++) {
    let type = Mock.mock({ 'type|1': ['DOMAIN', 'DOMAIN-SUFFIX', 'DOMAIN-KEYWORD', 'GEOIP'] });
    haveURL
      ? result.rules.push(type.type + ',' + Random.domain())
      : result.customRules.push(type.type + ',' + Random.domain());
  }
  return result;
}

function generateRules(req: Request, res: Response) {
  let result: RuleGroup[] = [];
  let count = Random.integer(5, 20);
  for (let i = 0; i < count; i++) {
    result.push(generateRule());
  }
  return res.json({
    status: 'success',
    data: result,
  });
}

function updateRule(req: Request, res: Response) {
  const { name, url } = req.body;
  const result = generateRule(name, url);
  result.lastUpdate = Random.now('second');
  setTimeout(() => {
    return res.json({
      status: 'success',
      data: result,
    });
  }, 1200);
}

function deleteRule(req: Request, res: Response) {
  if (Random.integer(0, 1)) {
    generateRules(req, res);
  } else {
    return res.json({
      status: 'fail',
      msg: 'error',
    });
  }
}

function addRule(req: Request, res: Response) {
  if (Random.integer(0, 1)) {
    generateRules(req, res);
  } else {
    return res.json({
      status: 'fail',
      msg: 'error',
    });
  }
}

export default {
  'GET /api/rules': generateRules,
  'POST /api/rule': (req: Request, res: Response) => {
    let action = req.query.action;
    action === 'add' && addRule(req, res);
    action === 'edit' && generateRules(req, res);
    action === 'update' && updateRule(req, res);
    action === 'delete' && deleteRule(req, res);
    action === 'updateall' && generateRules(req, res);
  },
};
