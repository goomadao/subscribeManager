import React from 'react';
import { connect } from 'dva';
import { Dispatch } from 'redux';
import Alert from '@material-ui/lab/Alert';
import { IconButton, Chip, Dialog, DialogTitle, DialogActions, Button } from '@material-ui/core';
import EditIcon from '@material-ui/icons/Edit';
import DeleteOutlineIcon from '@material-ui/icons/DeleteOutline';
import { Changer } from '../typing';
import { useSnackbar } from 'notistack';
import ChangerEditDialog from './ChangerEditDialog';

export interface ChangerProps {
  deleting: boolean;
  changer: Changer;
  dispatch: Dispatch;
  severty?: string;
}

function Changer({ changer, dispatch, severty, deleting }: ChangerProps) {
  const { enqueueSnackbar } = useSnackbar();

  const handleChangerDelete = () => {
    dispatch({
      type: 'changers/deleteChanger',
      payload: {
        emoji: changer.emoji,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(changer.emoji + '删除成功', { variant: 'success' });
          setDeleteDialogOpen(false);
        } else {
          enqueueSnackbar(msg || changer.emoji + '删除失败', { variant: 'error' });
        }
      },
    });
  };

  const handleChangerEdit = (name: string, changer: Changer) => {
    return dispatch({
      type: 'changers/editChanger',
      payload: {
        name,
        changer,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(changer.emoji + '编辑成功', { variant: 'success' });
          setEditDialogOpen(false);
        } else {
          enqueueSnackbar(msg || changer.emoji + '编辑失败', { variant: 'error' });
        }
      },
    });
  };

  const actions = [
    <IconButton onClick={() => setEditDialogOpen(true)}>
      <EditIcon fontSize="small" color="primary" />
    </IconButton>,
    <IconButton onClick={() => setDeleteDialogOpen(true)}>
      <DeleteOutlineIcon fontSize="small" color="secondary" />
    </IconButton>,
  ];
  const [editDialogOpen, setEditDialogOpen] = React.useState(false);
  const [deleteDialogOpen, setDeleteDialogOpen] = React.useState(false);
  const chip = (
    <div>
      {changer.regex &&
        changer.regex.split('|').map((str, i) => {
          switch (i % 3) {
            case 0:
              return <Chip label={str} color="default" style={{ margin: '3px' }} />;
            case 1:
              return <Chip label={str} color="primary" style={{ margin: '3px' }} />;
            case 2:
              return <Chip label={str} color="secondary" style={{ margin: '3px' }} />;
            default:
              return <Chip label={str} color="default" style={{ margin: '3px' }} />;
          }
        })}
    </div>
  );
  let alert: JSX.Element;
  switch (severty) {
    case 'error':
      alert = (
        <Alert icon={changer.emoji} action={actions} severity="error" style={{ margin: '10px' }}>
          {chip}
        </Alert>
      );
      break;
    case 'success':
      alert = (
        <Alert icon={changer.emoji} action={actions} severity="success" style={{ margin: '10px' }}>
          {chip}
        </Alert>
      );
      break;
    case 'info':
      alert = (
        <Alert icon={changer.emoji} action={actions} severity="info" style={{ margin: '10px' }}>
          {chip}
        </Alert>
      );
      break;
    case 'warning':
      alert = (
        <Alert icon={changer.emoji} action={actions} severity="warning" style={{ margin: '10px' }}>
          {chip}
        </Alert>
      );
      break;
    default:
      alert = (
        <Alert icon={changer.emoji} action={actions} severity="success" style={{ margin: '10px' }}>
          {chip}
        </Alert>
      );
  }
  return (
    <>
      {alert}
      <ChangerEditDialog
        open={editDialogOpen}
        changer={changer}
        dialogClose={() => setEditDialogOpen(false)}
        handleChangerEdit={handleChangerEdit}
      />
      <Dialog open={deleteDialogOpen} onClose={() => setDeleteDialogOpen(false)}>
        <DialogTitle>确定删除?</DialogTitle>
        <DialogActions>
          <Button color="primary" autoFocus={true} onClick={() => setDeleteDialogOpen(false)}>
            取消
          </Button>
          <Button color="secondary" disabled={deleting} onClick={handleChangerDelete}>
            删除
          </Button>
        </DialogActions>
        {/* <DisplayObject {...this.props.loading} /> */}
      </Dialog>
    </>
  );
}

export default connect(
  ({
    changers,
    loading,
  }: {
    changers: Changer[];
    loading: { effects: { [key: string]: boolean } };
  }) => ({
    changers,
    deleting: loading.effects['changers/deleteChanger'],
  }),
)(Changer);
