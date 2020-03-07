import request from '@/utils/request';
import { RuleGroup } from './typing';

export async function fetchRules() {
  return request('/api/rules');
}

export async function addRule(params: RuleGroup) {
  return request('/api/rule?action=add', {
    method: 'POST',
    data: params,
  });
}

export async function editRule(params: { name: string; rule: RuleGroup }) {
  return request('/api/rule?action=edit&rule=' + params.name, {
    method: 'POST',
    data: params.rule,
  });
}

export async function updateRule(params: { name: string }) {
  return request('/api/rule?action=update', {
    method: 'POST',
    data: params,
  });
}

export async function deleteRule(params: { name: string }) {
  return request('/api/rule?action=delete&rule=' + params.name, {
    method: 'POST',
    // data: params,
  });
}

export async function updateAllRules() {
  return request('/api/rule?action=updateall', { method: 'POST' });
}
