import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import { Node } from '../typing';
import {
  FormControl,
  InputLabel,
  FormHelperText,
  Select,
  MenuItem,
  Switch,
  FormControlLabel,
} from '@material-ui/core';
import { Formik } from 'formik';
import DisplayObject from '@/utils/DisplayObject';

export interface DialogProps {
  open: boolean;
  dialogClose: () => void;
  handleNodeAdd: (node: Node) => void;
  handleNodeEdit: (nodeName: string, node: Node) => void;
  node: Node | undefined;
}

class NodeEditDialog extends React.Component<DialogProps, object> {
  constructor(props: DialogProps) {
    super(props);
  }

  render() {
    const { open, dialogClose, node, handleNodeAdd, handleNodeEdit } = this.props;
    const initialValues: Node = node || { nodeType: 'ss' };
    return (
      <Dialog open={open} onClose={dialogClose} aria-labelledby="form-dialog-title">
        <Formik
          initialValues={initialValues}
          validate={values => {
            const errors: any = {};
            if (values.nodeType === 'ss' || values.nodeType === 'ssr') {
              if (!values.remarks) errors.remarks = 'Required';
              if (!values.server) errors.server = 'Required';
              if (!values.port) errors.port = 'Required';
              if (!values.password) errors.password = 'Required';
              if (!values.encryption) errors.encryption = 'Required';
            }
            if (values.nodeType === 'ssr') {
              if (!values.protocol) errors.protocol = 'Required';
              if (!values.obfs) errors.obfs = 'Required';
            }
            if (values.nodeType === 'vmess') {
              if (!values.ps) errors.ps = 'Required';
              if (!values.add) errors.add = 'Required';
              if (!values.port) errors.port = 'Required';
              if (!values.id) errors.id = 'Required';
              if (!values.aid) errors.aid = 'Required';
              if (!values.net) errors.net = 'Required';
            }
            return errors;
          }}
          onSubmit={values =>
            node ? handleNodeEdit(node.remarks || node.ps || '', values) : handleNodeAdd(values)
          }
        >
          {props => {
            const {
              values,
              errors,
              touched,
              isSubmitting,
              setErrors,
              setTouched,
              handleSubmit,
              handleChange,
              handleBlur,
            } = props;
            return (
              <>
                <DialogTitle id="form-dialog-title">
                  {node ? node.remarks || node.ps || 'Edit' : 'Add'}
                </DialogTitle>
                <form onSubmit={handleSubmit}>
                  <DialogContent>
                    <FormControl
                      required={true}
                      fullWidth={true}
                      error={errors.nodeType && touched.nodeType ? true : false}
                    >
                      <InputLabel htmlFor="nodeType">类型</InputLabel>
                      <Select
                        value={values.nodeType}
                        onBlur={handleBlur('nodeType')}
                        onChange={event => {
                          setErrors({});
                          setTouched({});
                          event.target.name = 'nodeType';
                          handleChange(event);
                        }}
                        id="nodeType"
                      >
                        <MenuItem value="ss">ss</MenuItem>
                        <MenuItem value="ssr">ssr</MenuItem>
                        <MenuItem value="vmess">vmess</MenuItem>
                      </Select>
                      {errors.nodeType && touched.nodeType && (
                        <FormHelperText>{errors.nodeType}</FormHelperText>
                      )}
                    </FormControl>
                    {values.nodeType === 'ss' && (
                      <>
                        <TextField
                          required={true}
                          fullWidth={true}
                          value={values.remarks}
                          error={errors.remarks && touched.remarks ? true : false}
                          helperText={errors.remarks && touched.remarks ? errors.remarks : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="名称"
                          id="remarks"
                        />
                        <TextField
                          required={true}
                          fullWidth={true}
                          value={values.server}
                          error={errors.server && touched.server ? true : false}
                          helperText={errors.server && touched.server ? errors.server : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="服务器IP"
                          id="server"
                        />
                        <TextField
                          required={true}
                          fullWidth={true}
                          type="number"
                          value={values.port}
                          error={errors.port && touched.port ? true : false}
                          helperText={errors.port && touched.port ? errors.port : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="端口"
                          id="port"
                        />
                        <TextField
                          required={true}
                          fullWidth={true}
                          value={values.password}
                          error={errors.password && touched.password ? true : false}
                          helperText={errors.password && touched.password ? errors.password : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="密码"
                          id="password"
                        />
                        <FormControl
                          required={true}
                          fullWidth={true}
                          error={errors.encryption && touched.encryption ? true : false}
                        >
                          <InputLabel htmlFor="encryption">加密方式</InputLabel>
                          <Select
                            value={values.encryption}
                            onBlur={handleBlur('encryption')}
                            onChange={handleChange('encryption')}
                            id="encryption"
                          >
                            <MenuItem value="aes-128-gcm">aes-128-gcm</MenuItem>
                            <MenuItem value="aes-192-gcm">aes-192-gcm</MenuItem>
                            <MenuItem value="aes-256-gcm">aes-256-gcm</MenuItem>
                            <MenuItem value="aes-128-cfb">aes-128-cfb</MenuItem>
                            <MenuItem value="aes-192-cfb">aes-192-cfb</MenuItem>
                            <MenuItem value="aes-256-cfb">aes-256-cfb</MenuItem>
                            <MenuItem value="aes-128-ctr">aes-128-ctr</MenuItem>
                            <MenuItem value="aes-192-ctr">aes-192-ctr</MenuItem>
                            <MenuItem value="aes-256-ctr">aes-256-ctr</MenuItem>
                            <MenuItem value="rc4-md5">rc4-md5</MenuItem>
                            <MenuItem value="chacha20-ietf">chacha20-ietf</MenuItem>
                            <MenuItem value="chacha20-ietf-poly1305">
                              chacha20-ietf-poly1305
                            </MenuItem>
                            <MenuItem value="xchacha20">xchacha20</MenuItem>
                            <MenuItem value="xchacha20-ietf-poly1305">
                              xchacha20-ietf-poly1305
                            </MenuItem>
                          </Select>
                          {errors.encryption && touched.encryption && (
                            <FormHelperText>{errors.encryption}</FormHelperText>
                          )}
                        </FormControl>
                        <TextField
                          fullWidth={true}
                          value={values.plugin}
                          error={errors.plugin && touched.plugin ? true : false}
                          helperText={errors.plugin && touched.plugin ? errors.plugin : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="插件"
                          id="plugin"
                        />
                        <TextField
                          fullWidth={true}
                          value={values.plugin_options}
                          error={errors.plugin_options && touched.plugin_options ? true : false}
                          helperText={
                            errors.plugin_options && touched.plugin_options
                              ? errors.plugin_options
                              : ''
                          }
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="插件选项"
                          id="plugin_options"
                        />
                      </>
                    )}
                    {values.nodeType === 'ssr' && (
                      <>
                        <TextField
                          required={true}
                          fullWidth={true}
                          value={values.remarks}
                          error={errors.remarks && touched.remarks ? true : false}
                          helperText={errors.remarks && touched.remarks ? errors.remarks : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="名称"
                          id="remarks"
                        />
                        <TextField
                          required={true}
                          fullWidth={true}
                          value={values.server}
                          error={errors.server && touched.server ? true : false}
                          helperText={errors.server && touched.server ? errors.server : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="服务器IP"
                          id="server"
                        />
                        <TextField
                          required={true}
                          fullWidth={true}
                          type="number"
                          value={values.port}
                          error={errors.port && touched.port ? true : false}
                          helperText={errors.port && touched.port ? errors.port : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="端口"
                          id="port"
                        />
                        <TextField
                          required={true}
                          fullWidth={true}
                          value={values.password}
                          error={errors.password && touched.password ? true : false}
                          helperText={errors.password && touched.password ? errors.password : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="密码"
                          id="password"
                        />
                        <FormControl
                          required={true}
                          fullWidth={true}
                          error={errors.encryption && touched.encryption ? true : false}
                        >
                          <InputLabel htmlFor="encryption">加密方式</InputLabel>
                          <Select
                            value={values.encryption}
                            onBlur={handleBlur('encryption')}
                            onChange={handleChange('encryption')}
                            id="encryption"
                          >
                            <MenuItem value="aes-128-gcm">aes-128-gcm</MenuItem>
                            <MenuItem value="aes-192-gcm">aes-192-gcm</MenuItem>
                            <MenuItem value="aes-256-gcm">aes-256-gcm</MenuItem>
                            <MenuItem value="aes-128-cfb">aes-128-cfb</MenuItem>
                            <MenuItem value="aes-192-cfb">aes-192-cfb</MenuItem>
                            <MenuItem value="aes-256-cfb">aes-256-cfb</MenuItem>
                            <MenuItem value="aes-128-ctr">aes-128-ctr</MenuItem>
                            <MenuItem value="aes-192-ctr">aes-192-ctr</MenuItem>
                            <MenuItem value="aes-256-ctr">aes-256-ctr</MenuItem>
                            <MenuItem value="rc4-md5">rc4-md5</MenuItem>
                            <MenuItem value="chacha20-ietf">chacha20-ietf</MenuItem>
                            <MenuItem value="chacha20-ietf-poly1305">
                              chacha20-ietf-poly1305
                            </MenuItem>
                            <MenuItem value="xchacha20">xchacha20</MenuItem>
                            <MenuItem value="xchacha20-ietf-poly1305">
                              xchacha20-ietf-poly1305
                            </MenuItem>
                          </Select>
                          {errors.encryption && touched.encryption && (
                            <FormHelperText>{errors.encryption}</FormHelperText>
                          )}
                        </FormControl>
                        <FormControl
                          required={true}
                          fullWidth={true}
                          error={errors.protocol && touched.protocol ? true : false}
                        >
                          <InputLabel htmlFor="protocol">协议</InputLabel>
                          <Select
                            value={values.protocol}
                            onBlur={handleBlur('protocol')}
                            onChange={handleChange('protocol')}
                            id="protocol"
                          >
                            <MenuItem value="origin">origin</MenuItem>
                            <MenuItem value="verify_deflate">verify_deflate</MenuItem>
                            <MenuItem value="auth_sha1_v4">auth_sha1_v4</MenuItem>
                            <MenuItem value="auth_aes128_md5">auth_aes128_md5</MenuItem>
                            <MenuItem value="auth_aes128_sha1">auth_aes128_sha1</MenuItem>
                            <MenuItem value="auth_chain_a">auth_chain_a</MenuItem>
                            <MenuItem value="auth_chain_b">auth_chain_b</MenuItem>
                          </Select>
                          {errors.protocol && touched.protocol && (
                            <FormHelperText>{errors.protocol}</FormHelperText>
                          )}
                        </FormControl>
                        <TextField
                          fullWidth={true}
                          value={values.protocol_param}
                          error={errors.protocol_param && touched.protocol_param ? true : false}
                          helperText={
                            errors.protocol_param && touched.protocol_param
                              ? errors.protocol_param
                              : ''
                          }
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="协议参数"
                          id="protocol_param"
                        />
                        <FormControl
                          required={true}
                          fullWidth={true}
                          error={errors.obfs && touched.obfs ? true : false}
                        >
                          <InputLabel htmlFor="obfs">混淆</InputLabel>
                          <Select
                            value={values.obfs}
                            onBlur={handleBlur('obfs')}
                            onChange={handleChange('obfs')}
                            id="obfs"
                          >
                            <MenuItem value="plain">plain</MenuItem>
                            <MenuItem value="http_simple">http_simple</MenuItem>
                            <MenuItem value="http_post">http_post</MenuItem>
                            <MenuItem value="random_head">random_head</MenuItem>
                            <MenuItem value="tls1.2_ticket_auth">tls1.2_ticket_auth</MenuItem>
                            <MenuItem value="tls1.2_ticket_fastauth">
                              tls1.2_ticket_fastauth
                            </MenuItem>
                          </Select>
                          {errors.obfs && touched.obfs && (
                            <FormHelperText>{errors.obfs}</FormHelperText>
                          )}
                        </FormControl>
                        <TextField
                          fullWidth={true}
                          value={values.obfs_param}
                          error={errors.obfs_param && touched.obfs_param ? true : false}
                          helperText={
                            errors.obfs_param && touched.obfs_param ? errors.obfs_param : ''
                          }
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="混淆参数"
                          id="obfs_param"
                        />
                      </>
                    )}
                    {values.nodeType === 'vmess' && (
                      <>
                        <TextField
                          required={true}
                          fullWidth={true}
                          value={values.ps}
                          error={errors.ps && touched.ps ? true : false}
                          helperText={errors.ps && touched.ps ? errors.ps : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="名称"
                          id="ps"
                        />
                        <TextField
                          required={true}
                          fullWidth={true}
                          value={values.add}
                          error={errors.add && touched.add ? true : false}
                          helperText={errors.add && touched.add ? errors.add : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="服务器IP"
                          id="add"
                        />
                        <TextField
                          required={true}
                          fullWidth={true}
                          type="number"
                          value={values.port}
                          error={errors.port && touched.port ? true : false}
                          helperText={errors.port && touched.port ? errors.port : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="端口"
                          id="port"
                        />
                        <TextField
                          required={true}
                          fullWidth={true}
                          value={values.id}
                          error={errors.id && touched.id ? true : false}
                          helperText={errors.id && touched.id ? errors.id : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="UUID"
                          id="id"
                        />
                        <TextField
                          required={true}
                          fullWidth={true}
                          type="number"
                          value={values.aid}
                          error={errors.aid && touched.aid ? true : false}
                          helperText={errors.aid && touched.aid ? errors.aid : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="AlterId"
                          id="aid"
                        />
                        <FormControl
                          required={true}
                          fullWidth={true}
                          error={errors.net && touched.net ? true : false}
                        >
                          <InputLabel htmlFor="net">Network</InputLabel>
                          <Select
                            value={values.net}
                            onBlur={handleBlur('net')}
                            onChange={handleChange('net')}
                            id="net"
                          >
                            <MenuItem value="ws">ws</MenuItem>
                            <MenuItem value="tcp">tcp</MenuItem>
                          </Select>
                          {errors.net && touched.net && (
                            <FormHelperText>{errors.net}</FormHelperText>
                          )}
                        </FormControl>
                        <TextField
                          fullWidth={true}
                          value={values.path}
                          error={errors.path && touched.path ? true : false}
                          helperText={errors.path && touched.path ? errors.path : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="Path"
                          id="path"
                        />
                        <TextField
                          fullWidth={true}
                          value={values.host}
                          error={errors.host && touched.host ? true : false}
                          helperText={errors.host && touched.host ? errors.host : ''}
                          onBlur={handleBlur}
                          onChange={handleChange}
                          label="Host"
                          id="host"
                        />
                        <FormControlLabel
                          value={values.tls === 'tls'}
                          onBlur={handleBlur('tls')}
                          onChange={handleChange('tls')}
                          control={<Switch color="primary" />}
                          label="TLS"
                          labelPlacement="start"
                          id="tls"
                        />
                      </>
                    )}
                    {/* <DisplayObject {...props} /> */}
                  </DialogContent>
                  <DialogActions>
                    <Button onClick={dialogClose} color="primary">
                      关闭
                    </Button>
                    {node ? (
                      <Button type="submit" disabled={isSubmitting} color="primary">
                        更新
                      </Button>
                    ) : (
                      <Button type="submit" disabled={isSubmitting} color="primary">
                        添加
                      </Button>
                    )}
                  </DialogActions>
                </form>
              </>
            );
          }}
        </Formik>
      </Dialog>
    );
  }
}

export default NodeEditDialog;
