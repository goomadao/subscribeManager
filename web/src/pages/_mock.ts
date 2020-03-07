import { Request, Response } from 'express';

function updateAll(req: Request, res: Response) {
  let num: number = Math.floor(Math.random() * 2);
  let result: any = num
    ? {}
    : {
        msg: 'error',
      };
  setTimeout(() => {
    return res.json(result);
  }, 1200);
}

export default {
  'POST /api/updateall': updateAll,
};
