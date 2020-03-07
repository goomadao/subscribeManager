import { Request, Response } from 'express';
import { Changer } from './typing';
import Mock from 'mockjs';

const Random = Mock.Random;

function generateChanger(emoji?: string) {
  let country = Mock.mock({
    'country|1': [
      '🇦🇱',
      '🇩🇿',
      '🇦🇫',
      '🇦🇷',
      '🇦🇪',
      '🇦🇼',
      '🇪🇬',
      '🇦🇴',
      '🇧🇯',
      '🇧🇪',
      '🇵🇪',
      '🇮🇸',
      '🇵🇷',
      '🇵🇱',
      '🇩🇰',
      '🇩🇪',
      '🇬🇬',
      '🇰🇷',
      '🇨🇦',
      '🇬🇦',
      '🇺🇸',
      '🇳🇴',
      '🇯🇵',
      '🇭🇰',
      '🇨🇳',
      '🇮🇳',
    ],
  });
  let result: Changer = {
    emoji: country.country,
    regex: '',
  };
  let count = Random.integer(10, 50);
  for (let i = 0; i < count; i++) {
    i === count - 1 ? (result.regex += Random.county()) : (result.regex += Random.county() + '|');
  }
  return result;
}

function generateChangers(req: Request, res: Response) {
  let result: Changer[] = [];
  let count = Random.integer(10, 50);
  for (let i = 0; i < count; i++) {
    result.push(generateChanger());
  }
  return res.json({
    status: 'success',
    data: result,
  });
}

function addChanger(req: Request, res: Response) {
  if (Random.integer(0, 1)) {
    generateChangers(req, res);
  } else {
    return res.json({
      status: 'fail',
      msg: 'error',
    });
  }
}

function deleteChanger(req: Request, res: Response) {
  if (Random.integer(0, 1)) {
    generateChangers(req, res);
  } else {
    return res.json({
      status: 'success',
      msg: 'error',
    });
  }
}

export default {
  'GET /api/changers': generateChangers,
  'POST /api/changer': (req: Request, res: Response) => {
    let action = req.query.action;
    action === 'add' && addChanger(req, res);
    action === 'delete' && deleteChanger(req, res);
    action === 'edit' && deleteChanger(req, res);
  },
};
