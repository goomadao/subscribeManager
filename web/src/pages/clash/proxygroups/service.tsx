import request from '@/utils/request';
import { Node, ClashProxyGroupSelector } from './typing';

export async function fetchGroups() {
  return request('/api/groups');
}

export async function fetchSelectors() {
  console.log('fetch selectors');
  return request('/api/selectors');
}

export async function addSelector(params: ClashProxyGroupSelector) {
  return request('/api/selector?action=add', {
    method: 'POST',
    data: params,
  });
}

export async function editSelector(params: { name: string; selector: ClashProxyGroupSelector }) {
  return request('/api/selector?action=edit&selector=' + params.name, {
    method: 'POST',
    data: params.selector,
  });
}

export async function updateSelector(params: any) {
  return request('/api/selector?action=update', {
    method: 'POST',
    data: {
      name: params.name,
      type: params.type,
    },
  });
}

export async function deleteSelector(params: { name: string }) {
  return request('/api/selector?action=delete&selector=' + params.name, { method: 'POST' });
}

export async function updateAllSelectors() {
  return request('/api/selector?action=updateall', { method: 'POST' });
}
