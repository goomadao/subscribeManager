import React from 'react';
import { connect, Dispatch } from 'dva';
import NodeGroup from './components/NodeGroup';
import { Group, Node } from './typing';
import { Button, Typography, Backdrop, CircularProgress } from '@material-ui/core';
import GroupAddDialog from './components/GroupAddDialog';
import { withSnackbar, WithSnackbarProps } from 'notistack';
import styles from './index.css';
import AddCircleOutlineIcon from '@material-ui/icons/AddCircleOutline';
import SyncIcon from '@material-ui/icons/Sync';
import NodePanel from '@/components/NodePanel';

interface NodeGroupsProps extends WithSnackbarProps {
  updatingAll: boolean;
  groups: Group[];
  dispatch: Dispatch;
}

interface NodeGroupsState {
  groupAddDialogOpen: boolean;
  nodeDisplay: Node;
  nodePanelOpen: boolean;
}

class NodeGroups extends React.Component<NodeGroupsProps, NodeGroupsState> {
  constructor(props: NodeGroupsProps) {
    super(props);
    this.state = {
      groupAddDialogOpen: false,
      nodeDisplay: { nodeType: '' },
      nodePanelOpen: false,
    };
  }

  componentWillMount() {
    const { dispatch } = this.props;
    dispatch({
      type: 'groups/fetchGroups',
    });
  }

  handleGroupAddDialogClose = () => {
    this.setState({ groupAddDialogOpen: false });
  };

  handleGroupAdd = (group: Group) => {
    const { dispatch, enqueueSnackbar } = this.props;
    return dispatch({
      type: 'groups/addGroup',
      payload: {
        name: group.name,
        url: group.url,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(group.name + '添加成功', { variant: 'success' });
          this.handleGroupAddDialogClose();
        } else {
          enqueueSnackbar(msg || group.name + '添加失败', { variant: 'error' });
        }
      },
    });
  };

  handleAllGroupsUpdate = () => {
    const { dispatch, enqueueSnackbar } = this.props;
    dispatch({
      type: 'groups/updateAllGroups',
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar('节点组更新成功', { variant: 'success' });
          this.handleGroupAddDialogClose();
        } else {
          enqueueSnackbar(msg || '节点组更新失败', { variant: 'error' });
        }
      },
    });
  };

  handleNodeClick = (node: Node) => {
    this.setState({ nodeDisplay: node, nodePanelOpen: true });
  };

  handleNodePanelClose = () => {
    this.setState({ nodeDisplay: { nodeType: '' }, nodePanelOpen: false });
  };

  render() {
    const { groups, updatingAll } = this.props;
    const { groupAddDialogOpen, nodeDisplay, nodePanelOpen } = this.state;
    return (
      <>
        <GroupAddDialog
          open={groupAddDialogOpen}
          dialogClose={this.handleGroupAddDialogClose}
          handleGroupAdd={this.handleGroupAdd}
        />
        <NodePanel
          node={nodeDisplay}
          open={nodePanelOpen}
          handleNodePanelClose={this.handleNodePanelClose}
        />
        <div className={styles.globalButtons}>
          <Button
            variant="contained"
            color="secondary"
            startIcon={<SyncIcon />}
            disabled={updatingAll || !groups.length}
            onClick={this.handleAllGroupsUpdate}
          >
            刷新所有节点组
          </Button>
          <Button
            variant="contained"
            color="primary"
            endIcon={<AddCircleOutlineIcon />}
            onClick={() => {
              this.setState({ groupAddDialogOpen: true });
            }}
          >
            添加节点组
          </Button>
        </div>
        {Array.isArray(groups) && groups.length ? (
          groups.map(group => (
            <NodeGroup
              group={group}
              updatingAll={updatingAll}
              handleNodeClick={this.handleNodeClick}
            />
          ))
        ) : (
          <Typography>还没有节点组，点击上方按钮添加。</Typography>
        )}
      </>
    );
  }
}

export default withSnackbar(
  connect(
    ({
      groups,
      loading,
    }: {
      groups: Group[];
      loading: {
        effects: { [key: string]: boolean };
      };
    }) => ({
      groups,
      adding: loading.effects['groups/addGroup'],
      updatingAll: loading.effects['groups/updateAllGroups'],
    }),
  )(NodeGroups),
);
