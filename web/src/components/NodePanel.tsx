import React from 'react';
import { makeStyles, Theme } from '@material-ui/core/styles';
import {
  AppBar,
  Tabs,
  Tab,
  Typography,
  Box,
  Dialog,
  Button,
  Card,
  CardHeader,
  CardContent,
  Divider,
} from '@material-ui/core';
import { Node } from '@/pages/groups/typing';
import SyntaxHighlighter from 'react-syntax-highlighter';
import { docco, vs2015 } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import CopyToClipboard from 'react-copy-to-clipboard';
import QRCode from 'qrcode.react';
import { withSnackbar, WithSnackbarProps } from 'notistack';

interface PanelProps extends WithSnackbarProps {
  node: Node;
  open: boolean;
  handleNodePanelClose: () => void;
}

interface PanelState {
  tabValue: number;
}

interface SSConfig {
  server: string;
  local_address: string;
  local_port: number;
  timeout: number;
  workers: number;
  server_port: number;
  password: string;
  method: string;
  plugin?: string;
  plugin_options?: string;
  remarks: string;
}

interface SSRConfig {
  server: string;
  local_address: string;
  local_port: number;
  timeout: number;
  workers: number;
  server_port: number;
  password: string;
  method: string;
  obfs: string;
  obfs_param: string;
  protocol: string;
  protocol_param: string;
  remarks: string;
}

interface VmessConfig {
  v: string;
  ps: string;
  add: string;
  port: string;
  id: string;
  aid: string;
  net: string;
  type: string;
  host: string;
  path: string;
  tls: string;
}

interface HTTPConfig {
  server: string;
  name: string;
  port: number;
  username: string;
  password: string;
  tls: string;
  skipCertVerify: boolean;
}

interface Socks5Config {
  server: string;
  name: string;
  port: number;
  username: string;
  password: string;
  tls: string;
  skipCertVerify: boolean;
}

function a11yProps(index: any) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

const useStyles = makeStyles((theme: Theme) => ({
  root: {
    flexGrow: 1,
    backgroundColor: theme.palette.background.paper,
    display: 'flex',
  },
}));

class NodePanel extends React.Component<PanelProps, PanelState> {
  constructor(props: PanelProps) {
    super(props);
    this.state = {
      tabValue: 0,
    };
  }

  handleChange = (event: React.ChangeEvent<{}>, newValue: number) => {
    this.setState({ tabValue: newValue });
  };

  renderSSJson = () => {
    const { node } = this.props;
    let ss: SSConfig = {
      server: node.server || '',
      local_address: '127.0.0.1',
      local_port: 1080,
      timeout: 300,
      workers: 1,
      server_port: node.port || 0,
      password: node.password || '',
      method: node.encryption || '',
      plugin: node.plugin,
      plugin_options: node.plugin_options?.toString(),
      remarks: node.remarks || '',
    };
    return ss;
  };

  renderSSRJson = () => {
    const { node } = this.props;
    let ssr: SSRConfig = {
      server: node.server || '',
      local_address: '127.0.0.1',
      local_port: 1080,
      timeout: 300,
      workers: 1,
      server_port: node.port || 0,
      password: node.password || '',
      method: node.encryption || '',
      obfs: node.obfs || '',
      obfs_param: node.obfs_param || '',
      protocol: node.protocol || '',
      protocol_param: node.protocol_param || '',
      remarks: node.remarks || '',
    };
    return ssr;
  };

  renderVmessJson = () => {
    const { node } = this.props;
    let vmess: VmessConfig = {
      v: node.v || '2',
      ps: node.ps || '',
      add: node.add || '',
      port: node.port?.toString() || '0',
      id: node.id || '',
      aid: node.aid?.toString() || '100',
      net: node.net || '',
      type: node.type || '',
      host: node.host || '',
      path: node.path || '',
      tls: node.tls || '',
    };
    return vmess;
  };

  renderHTTPJson = () => {
    const { node } = this.props;
    let http: HTTPConfig = {
      server: node.server || '',
      name: node.remarks || '',
      port: node.port || 0,
      username: node.username || '',
      password: node.password || '',
      tls: node.tls || '',
      skipCertVerify: node.skipCertVerify || false,
    };
    return http;
  };

  renderSocks5Json = () => {
    const { node } = this.props;
    let socks5: Socks5Config = {
      server: node.server || '',
      name: node.remarks || '',
      port: node.port || 0,
      username: node.username || '',
      password: node.password || '',
      tls: node.tls || '',
      skipCertVerify: node.skipCertVerify || false,
    };
    return socks5;
  };

  renderJson = () => {
    const { node } = this.props;
    let res: any = {};
    switch (node.nodeType) {
      case 'ss':
        res = this.renderSSJson();
        break;
      case 'ssr':
        res = this.renderSSRJson();
        break;
      case 'vmess':
        res = this.renderVmessJson();
        break;
      case 'http':
        res = this.renderHTTPJson();
        break;
      case 'socks5':
        res = this.renderSocks5Json();
        break;
      default:
        return <div>error</div>;
    }
    return (
      <SyntaxHighlighter language="json" style={vs2015} showLineNumbers={true}>
        {JSON.stringify(res, null, 2)}
      </SyntaxHighlighter>
    );
  };

  renderSSLink = () => {
    const { node } = this.props;
    let ss = 'ss://';
    let userinfo = btoa(node.encryption + ':' + node.password);
    ss = ss + userinfo + '@' + node.server + ':' + node.port?.toString();
    let plugin = '';
    if (node.plugin) {
      plugin = '/?plugin=' + encodeURI(node.plugin + ';' + node.plugin_options?.toString());
    }
    return ss + plugin + (node.remarks ? '#' + encodeURI(node.remarks) : '');
  };

  renderSSRLink = () => {
    const { node } = this.props;
    let link: string =
      node.server +
      ':' +
      node.port?.toString +
      ':' +
      node.protocol +
      ':' +
      node.encryption +
      ':' +
      node.obfs +
      ':' +
      btoa(node.password || '') +
      '/obfsparam=' +
      btoa(node.obfs_param || '') +
      '&protoparam=' +
      btoa(node.protocol_param || '') +
      '&remarks=' +
      btoa(unescape(encodeURIComponent(node.remarks || ''))) +
      '&group=' +
      btoa(unescape(encodeURIComponent(node.group || '')));
    return 'ssr://' + btoa(link);
  };

  renderVmessLink = () => {
    let link: string = JSON.stringify(this.renderVmessJson());
    return 'vmess://' + btoa(unescape(encodeURIComponent(link)));
  };

  renderLink = () => {
    const { node } = this.props;
    let link: string = '';
    switch (node.nodeType) {
      case 'ss':
        link = this.renderSSLink();
        break;
      case 'ssr':
        link = this.renderSSRLink();
        break;
      case 'vmess':
        link = this.renderVmessLink();
        break;
      default:
        return <div>error</div>;
    }
    return (
      <>
        <p>{link}</p>
        <CopyToClipboard
          text={link}
          onCopy={() =>
            this.props.enqueueSnackbar('复制成功', {
              variant: 'success',
              preventDuplicate: true,
            })
          }
        >
          <Button color="primary">点击复制</Button>
        </CopyToClipboard>
        <Button href={link} color="primary">
          点击自动导入
        </Button>
      </>
    );
  };

  renderQRCode = () => {
    const { node } = this.props;
    let link: string = '';
    if (node.nodeType === 'ss') link = this.renderSSLink();
    else if (node.nodeType === 'ssr') link = this.renderSSRLink();
    else if (node.nodeType === 'vmess') link = this.renderVmessLink();
    else return <div>error</div>;
    return <QRCode value={link} />;
  };

  renderLinkAndQRCode = () => {
    const { node } = this.props;
    if (node.nodeType === 'http' || node.nodeType === 'socks5') {
      return false;
    }
    return true;
  };

  render() {
    const { open, handleNodePanelClose } = this.props;
    const { tabValue } = this.state;
    return (
      <Dialog open={open} onClose={handleNodePanelClose} fullWidth={true}>
        {/* <div className={this.classes.root}> */}
        {/* <div style={{ height: 480 }}> */}
        {/* <AppBar position="static"> */}
        {/* <Tabs
            value={tabValue}
            onChange={this.handleChange}
            variant="fullWidth"
            // orientation="vertical"
            aria-label="simple tabs example"
          >
            <Tab label="Json" {...a11yProps(0)} />
            <Tab label="链接" {...a11yProps(1)} />
            <Tab label="二维码" {...a11yProps(2)} />
          </Tabs> */}
        {/* </AppBar> */}
        {/* {tabValue === 0 && this.renderJson()}
          {tabValue === 1 && this.renderLink()}
          {tabValue === 2 && (
            <div style={{ display: 'flex', justifyContent: 'center' }}>{this.renderQRCode()}</div>
          )} */}
        <div style={{ display: 'flex', flexDirection: 'column' }}>
          <Card style={{ margin: '10px' }}>
            <CardHeader title="JSON" />
            <Divider />
            <CardContent>{this.renderJson()}</CardContent>
          </Card>
          {this.renderLinkAndQRCode() && (
            <Card style={{ margin: '10px' }}>
              <CardHeader title="链接" />
              <Divider />
              <CardContent>{this.renderLink()}</CardContent>
            </Card>
          )}
          {this.renderLinkAndQRCode() && (
            <Card style={{ margin: '10px' }}>
              <CardHeader title="二维码" />
              <Divider />
              <CardContent>
                <div style={{ display: 'flex', justifyContent: 'center' }}>
                  {this.renderQRCode()}
                </div>
              </CardContent>
            </Card>
          )}
        </div>
        {/* </div> */}
      </Dialog>
    );
  }
}

export default withSnackbar(NodePanel);
