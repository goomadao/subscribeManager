import React from 'react';
import { connect, Dispatch } from 'dva';
import { Changer } from './typing';
import { Button, Typography } from '@material-ui/core';
import ChangerAddDialog from './components/ChangerAddDialog';
import { withSnackbar, WithSnackbarProps } from 'notistack';
import styles from './index.css';
import AddCircleOutlineIcon from '@material-ui/icons/AddCircleOutline';
import ChangerAlert from './components/ChangerAlert';

interface NodeGroupsProps extends WithSnackbarProps {
  changers: Changer[];
  dispatch: Dispatch;
}

interface NodeGroupsState {
  changerAddDialogOpen: boolean;
}

class NodeGroups extends React.Component<NodeGroupsProps, NodeGroupsState> {
  constructor(props: NodeGroupsProps) {
    super(props);
    this.state = {
      changerAddDialogOpen: false,
    };
  }

  componentWillMount() {
    const { dispatch } = this.props;
    dispatch({
      type: 'changers/fetchChangers',
    });
  }

  handleChangerAddDialogClose = () => {
    this.setState({ changerAddDialogOpen: false });
  };

  handleChangerAdd = (changer: Changer) => {
    const { dispatch, enqueueSnackbar } = this.props;
    return dispatch({
      type: 'changers/addChanger',
      payload: changer,
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(changer.emoji + '添加成功', { variant: 'success' });
          this.handleChangerAddDialogClose();
        } else {
          enqueueSnackbar(msg || changer.emoji + '添加失败', { variant: 'error' });
        }
      },
    });
  };

  render() {
    const { changers } = this.props;
    const { changerAddDialogOpen } = this.state;
    let alertSeverty = ['error', 'success', 'info', 'warning']; //red,green,blue,orange 500
    return (
      <>
        <ChangerAddDialog
          open={changerAddDialogOpen}
          dialogClose={this.handleChangerAddDialogClose}
          handleChangerAdd={this.handleChangerAdd}
        />
        <div className={styles.globalButtons}>
          <Button
            variant="contained"
            color="primary"
            endIcon={<AddCircleOutlineIcon />}
            onClick={() => {
              this.setState({ changerAddDialogOpen: true });
            }}
          >
            添加emoji替换规则
          </Button>
        </div>
        {Array.isArray(changers) && changers.length ? (
          changers.map((changer, i) => (
            <ChangerAlert changer={changer} severty={alertSeverty[i % 4]} />
          ))
        ) : (
          <Typography>还没有自定义emoji，点击上方按钮添加.</Typography>
        )}
        {/* <div><DisplayObject {...this.props.loading} /></div> */}
      </>
    );
  }
}

export default withSnackbar(
  connect(
    ({
      changers,
      loading,
    }: {
      changers: Changer[];
      loading: {
        effects: { [key: string]: boolean };
      };
    }) => ({
      changers,
    }),
  )(NodeGroups),
);
