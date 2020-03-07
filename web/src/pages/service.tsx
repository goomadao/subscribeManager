import request from '@/utils/request';

export async function updateAll() {
  return request('/api/updateall', {
    method: 'POST',
  });
}
