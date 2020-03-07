import React from 'react';
import { Dispatch } from 'redux';
import { connect } from 'dva';
import ProxyGroup from './components/ProxyGroup';
import { ClashProxyGroupSelector, Group, Node } from './typing';
import { Button, Typography, Backdrop, CircularProgress } from '@material-ui/core';
import SelectorAddDialog from './components/SelectorAddDialog';
import { withSnackbar, WithSnackbarProps } from 'notistack';
import styles from './index.css';
import AddCircleOutlineIcon from '@material-ui/icons/AddCircleOutline';
import SyncIcon from '@material-ui/icons/Sync';
import NodePanel from '@/components/NodePanel';

interface NodeGroupsProps extends WithSnackbarProps {
  updatingAll: boolean;
  selectors: ClashProxyGroupSelector[];
  groups: Group[];
  dispatch: Dispatch;
}

interface NodeGroupsState {
  selectorAddDialogOpen: boolean;
  nodeDisplay: Node;
  nodePanelOpen: boolean;
}

class NodeGroups extends React.Component<NodeGroupsProps, NodeGroupsState> {
  constructor(props: NodeGroupsProps) {
    super(props);
    this.state = {
      selectorAddDialogOpen: false,
      nodeDisplay: { nodeType: '' },
      nodePanelOpen: false,
    };
  }

  componentWillMount() {
    const { dispatch } = this.props;
    dispatch({
      type: 'selectors/fetchSelectors',
    });
    dispatch({ type: 'selectors/fetchGroups' });
  }

  handleSelectorAddDialogClose = () => {
    this.setState({ selectorAddDialogOpen: false });
  };

  handleSelectorAdd = (selector: ClashProxyGroupSelector) => {
    const { dispatch, enqueueSnackbar } = this.props;
    return dispatch({
      type: 'selectors/addSelector',
      payload: selector,
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(selector.name + '添加成功', { variant: 'success' });
          this.handleSelectorAddDialogClose();
        } else {
          enqueueSnackbar(msg || selector.name + '添加失败', { variant: 'error' });
        }
      },
    });
  };

  handleAllSelectorsUpdate = () => {
    const { dispatch, enqueueSnackbar } = this.props;
    dispatch({
      type: 'selectors/updateAllSelectors',
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar('代理组更新成功', { variant: 'success' });
          this.handleSelectorAddDialogClose();
        } else {
          enqueueSnackbar(msg || '代理组更新失败', { variant: 'error' });
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
    const { selectors, groups, updatingAll } = this.props;
    const { selectorAddDialogOpen, nodeDisplay, nodePanelOpen } = this.state;
    return (
      <>
        <SelectorAddDialog
          open={selectorAddDialogOpen}
          selectors={selectors}
          groups={groups}
          dialogClose={this.handleSelectorAddDialogClose}
          handleSelectorAdd={this.handleSelectorAdd}
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
            disabled={updatingAll || !selectors?.length}
            onClick={this.handleAllSelectorsUpdate}
          >
            刷新所有代理组
          </Button>
          <Button
            variant="contained"
            color="primary"
            endIcon={<AddCircleOutlineIcon />}
            onClick={() => {
              this.setState({ selectorAddDialogOpen: true });
            }}
          >
            添加代理组
          </Button>
        </div>
        {Array.isArray(selectors) && selectors.length ? (
          selectors.map(selector => (
            <ProxyGroup
              selector={selector}
              updatingAll={updatingAll}
              handleNodeClick={this.handleNodeClick}
            />
          ))
        ) : (
          <Typography>还没有代理组，点击上方按钮添加。</Typography>
        )}
      </>
    );
  }
}

export default withSnackbar(
  connect(
    ({
      selectors,
      loading,
    }: {
      selectors: { selectors: ClashProxyGroupSelector[]; groups: Group[] };
      loading: { effects: { [key: string]: boolean } };
    }) => ({
      selectors: selectors.selectors,
      groups: selectors.groups,
      adding: loading.effects['relectoes/addGroup'],
      updatingAll: loading.effects['selectors/updateAllSelectors'],
    }),
  )(NodeGroups),
);
