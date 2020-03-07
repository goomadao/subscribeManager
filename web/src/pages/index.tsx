import React from 'react';
import QRCode from 'qrcode.react';
import CopyToClipboard from 'react-copy-to-clipboard';
import styles from './index.css';
import {
  Card,
  CardHeader,
  Avatar,
  CardContent,
  Button,
  Divider,
  Snackbar,
  CircularProgress,
} from '@material-ui/core';
import DoneIcon from '@material-ui/icons/Done';
import { green } from '@material-ui/core/colors';
import MuiAlert, { AlertProps } from '@material-ui/lab/Alert';
import { useSnackbar, withSnackbar, WithSnackbarProps } from 'notistack';
import { connect, Dispatch } from 'dva';

function Alert(props: AlertProps) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

interface Props {
  api: string;
  title: string;
  directImport?: string;
}

interface IndexProps extends WithSnackbarProps {
  updating: boolean;
  dispatch: Dispatch;
}

const Subscription = (props: Props) => {
  const { enqueueSnackbar } = useSnackbar();

  const { directImport, api, title } = props;

  const url = window.location.protocol + '//' + window.location.host + api;

  return (
    <Card className={styles.card}>
      <CardHeader avatar={<Avatar>{title}</Avatar>} title={title} />
      <CardContent>
        <div className={styles.content}>
          <QRCode value={url} size={200} />
          <Divider orientation="vertical" />
          <div className={styles.buttons}>
            {directImport && (
              <Button color="secondary" href={directImport}>
                一键导入订阅
              </Button>
            )}
            <CopyToClipboard
              text={url}
              onCopy={() => {
                enqueueSnackbar(title + '订阅链接复制成功', {
                  variant: 'success',
                  preventDuplicate: true,
                });
              }}
            >
              <Button color="primary">复制订阅链接</Button>
            </CopyToClipboard>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default withSnackbar(
  connect(({ loading }: { loading: { effects: { [key: string]: boolean } } }) => ({
    updating: loading.effects['subs/updateAll'],
  }))(function(props: IndexProps) {
    const handleUpdateAll = () => {
      const { dispatch, enqueueSnackbar } = props;
      setJustFinish(true);
      dispatch({
        type: 'subs/updateAll',
        callback: (success: boolean, msg?: string) => {
          if (!success) {
            enqueueSnackbar(msg || '更新出错', { variant: 'error' });
          }
          setTimeout(() => {
            setJustFinish(false);
          }, 2000);
        },
      });
    };

    const [justFinish, setJustFinish] = React.useState(false);

    return (
      <>
        <div style={{ display: 'flex', justifyContent: 'center' }}>
          <Button
            color="secondary"
            variant="contained"
            onClick={handleUpdateAll}
            disabled={props.updating}
          >
            刷新所有节点组、规则组和代理组
          </Button>
          {props.updating ? (
            <CircularProgress />
          ) : justFinish ? (
            <DoneIcon style={{ color: green[500] }} fontSize="large" />
          ) : null}
        </div>
        <div className={styles.subscription}>
          <Subscription
            title="ClashR"
            api="/api/sub?class=clashr"
            directImport={
              'clash://install-config?url=' +
              encodeURI(
                window.location.protocol + '//' + window.location.host + '/api/sub?class=clashr',
              )
            }
          />
          {/* <Subscription title="SS" api="/api/sub?class=ss" />
          <Subscription title="SSR" api="/api/sub?clash=ssr" />
          <Subscription title="SSD" api="/api/sub?class=ssd" /> */}
        </div>
      </>
    );
  }),
);
