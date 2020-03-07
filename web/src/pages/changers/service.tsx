import request from '@/utils/request';
import { Changer } from './typing';

export async function fetchChangers() {
  return request('/api/changers');
}

export async function addChanger(params: Changer) {
  return request('/api/changer?action=add', {
    method: 'POST',
    data: params,
  });
}

export async function editChanger(params: { name: string; changer: Changer }) {
  return request('/api/changer?action=edit&changer=' + params.name, {
    method: 'POST',
    data: params.changer,
  });
}

export async function deleteChanger(params: { emoji: string }) {
  return request('/api/changer?action=delete&changer=' + params.emoji, {
    method: 'POST',
    // data: params,
  });
}
