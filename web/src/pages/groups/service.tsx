import request from '@/utils/request';
import { Node, Group } from './typing';

export async function fetchGroups() {
  return request('/api/groups');
}

export async function addGroup(params: { name: string; url: string }) {
  return request('/api/group?action=add', {
    method: 'POST',
    data: params,
  });
}

export async function editGroup(params: { name: string; group: Group }) {
  return request('/api/group?action=edit&group=' + params.name, {
    method: 'POST',
    data: params.group,
  });
}

export async function updateGroup(params: any) {
  return request('/api/group?action=update', {
    method: 'POST',
    data: {
      name: params.name,
    },
  });
}

export async function deleteGroup(params: { name: string }) {
  return request('/api/group?action=delete&group=' + params.name, { method: 'POST' });
}

export async function updateAllGroups() {
  return request('/api/group?action=updateall', { method: 'POST' });
}

export async function addNode(params: { groupName: string; node: Node }) {
  return request(
    '/api/node?action=add&group=' + params.groupName + '&type=' + params.node.nodeType,
    {
      method: 'POST',
      data: params.node,
    },
  );
}

export async function editNode(params: { groupName: string; nodeName: string; node: Node }) {
  return request(
    '/api/node?action=edit&group=' +
      params.groupName +
      '&node=' +
      params.nodeName +
      '&type=' +
      params.node.nodeType,
    {
      method: 'POST',
      data: params.node,
    },
  );
}

export async function deleteNode(params: { groupName: string; nodeName: string }) {
  return request('/api/node?action=delete&group=' + params.groupName + '&node=' + params.nodeName, {
    method: 'POST',
  });
}
