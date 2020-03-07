import React from 'react';
import styles from './NodeGroup.css';
import { makeStyles, withStyles } from '@material-ui/core/styles';
import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import MuiExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import Typography from '@material-ui/core/Typography';
import AddIcon from '@material-ui/icons/Add';
import RefreshIcon from '@material-ui/icons/Refresh';
import EditIcon from '@material-ui/icons/Edit';
import DeleteIcon from '@material-ui/icons/Delete';
import {
  Button,
  IconButton,
  List,
  ListItem,
  CircularProgress,
  ListItemText,
  ListItemSecondaryAction,
  Dialog,
  DialogTitle,
  DialogActions,
  Chip,
  Tooltip,
} from '@material-ui/core';
import { Node, Group } from '../typing';
import { connect, Dispatch } from 'dva';
import NodeEditDialog from './NodeEditDialog';
import { withSnackbar, WithSnackbarProps } from 'notistack';
import GroupEditDialog from './GroupEditDialog';

const ExpansionPanelSummary = withStyles({
  content: {
    justifyContent: 'space-between',
  },
})(MuiExpansionPanelSummary);

export interface GroupProps extends WithSnackbarProps {
  group: Group;
  updatingAll: boolean;
  dispatch: Dispatch;
  nodeDeleting: boolean;
  groupDeleting: boolean;
  handleNodeClick: (node: Node) => void;
}

export interface GroupState {
  updating: boolean;
  nodeEditDialogOpen: boolean;
  nodeDeleteDialogOpen: boolean;
  groupAddDialogOpen: boolean;
  groupEditDialogOpen: boolean;
  groupDeleteDialogOpen: boolean;
  editedNode?: Node;
  deletedNode?: Node;
}

class NodeGroup extends React.Component<GroupProps, GroupState> {
  constructor(props: GroupProps) {
    super(props);
    this.state = {
      updating: false,
      nodeEditDialogOpen: false,
      nodeDeleteDialogOpen: false,
      groupDeleteDialogOpen: false,
      groupAddDialogOpen: false,
      groupEditDialogOpen: false,
    };
  }

  handleGroupUpdate = (event: any) => {
    const { group, dispatch, enqueueSnackbar } = this.props;
    event.stopPropagation();
    this.setState({
      updating: true,
    });
    dispatch({
      type: 'groups/updateGroup',
      payload: {
        name: group.name,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(group.name + '更新成功', {
            variant: 'success',
          });
        } else {
          enqueueSnackbar(msg || group.name + '更新失败', {
            variant: 'error',
          });
        }
        this.setState({
          updating: false,
        });
      },
    });
  };

  handleGroupEdit = (newGroup: Group) => {
    const { group, dispatch, enqueueSnackbar } = this.props;
    return dispatch({
      type: 'groups/editGroup',
      payload: {
        name: group.name,
        group: newGroup,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(group.name + '修改成功', { variant: 'success' });
          this.groupEditDialogClose();
        } else {
          enqueueSnackbar(msg || group.name + '修改失败', { variant: 'error' });
        }
      },
    });
  };

  handleGroupDelete = (e: any) => {
    e.stopPropagation();
    const { group, dispatch, enqueueSnackbar } = this.props;
    dispatch({
      type: 'groups/deleteGroup',
      payload: {
        name: group.name,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(group.name + '删除成功', {
            variant: 'success',
          });
          this.setState({
            groupDeleteDialogOpen: false,
          });
        } else {
          enqueueSnackbar(msg || group.name + '删除失败', {
            variant: 'error',
          });
        }
      },
    });
  };

  handleNodeAdd = (node: Node) => {
    const { group, dispatch, enqueueSnackbar } = this.props;
    return dispatch({
      type: 'groups/addNode',
      payload: {
        groupName: group.name,
        node: node,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar('添加节点成功', {
            variant: 'success',
          });
          this.nodeEditDialogClose();
        } else {
          enqueueSnackbar(msg || '添加节点失败', {
            variant: 'error',
          });
        }
      },
    });
  };

  handleNodeEdit = (name: string, node: Node) => {
    const { group, dispatch, enqueueSnackbar } = this.props;
    return dispatch({
      type: 'groups/editNode',
      payload: {
        groupName: group.name,
        nodeName: name,
        node: node,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar('修改节点成功', {
            variant: 'success',
          });
          this.nodeEditDialogClose();
        } else {
          enqueueSnackbar(msg || '修改节点失败', {
            variant: 'error',
          });
        }
      },
    });
  };

  handleNodeDelete = (e: any) => {
    e.stopPropagation();
    const { group, dispatch, enqueueSnackbar } = this.props;
    const { deletedNode } = this.state;
    dispatch({
      type: 'groups/deleteNode',
      payload: {
        groupName: group.name,
        nodeName: deletedNode?.remarks || deletedNode?.ps,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar((deletedNode?.remarks || deletedNode?.ps) + '删除成功', {
            variant: 'success',
          });
          this.setState({
            nodeDeleteDialogOpen: false,
          });
        } else {
          enqueueSnackbar(msg || (deletedNode?.remarks || deletedNode?.ps) + '删除失败', {
            variant: 'error',
          });
        }
      },
    });
  };

  nodeEditDialogOpen = (event: any, node?: Node) => {
    event.stopPropagation();
    this.setState({
      nodeEditDialogOpen: true,
      editedNode: node,
    });
  };

  groupEditDialogOpen = (e: any) => {
    e.stopPropagation();
    this.setState({ groupEditDialogOpen: true });
  };

  nodeEditDialogClose = () => {
    this.setState({
      nodeEditDialogOpen: false,
    });
  };

  groupDeleteDialogClose = () => {
    this.setState({ groupDeleteDialogOpen: false });
  };

  groupEditDialogClose = () => {
    this.setState({ groupEditDialogOpen: false });
  };

  nodeDeleteDialogClose = () => {
    this.setState({ nodeDeleteDialogOpen: false });
  };

  render() {
    const { group, nodeDeleting, groupDeleting, updatingAll, handleNodeClick } = this.props;
    const {
      updating,
      nodeEditDialogOpen,
      editedNode,
      nodeDeleteDialogOpen,
      groupDeleteDialogOpen,
      groupEditDialogOpen,
    } = this.state;
    return (
      <ExpansionPanel>
        <ExpansionPanelSummary
          aria-label="Expand"
          aria-controls="additional-actions1-content"
          id="additional-actions1-header"
        >
          <Typography color="textPrimary">
            {group.name}
            {group.url && (
              <Chip
                label={group.url}
                color="secondary"
                variant="outlined"
                style={{ marginLeft: '10px' }}
              />
            )}
          </Typography>
          <Typography color="textSecondary">
            <div className={styles.groupButtons}>
              <Tooltip title="添加节点">
                <IconButton
                  onClick={this.nodeEditDialogOpen}
                  onFocus={event => {
                    event.stopPropagation();
                  }}
                  disabled={group.url.length > 0}
                >
                  <AddIcon
                    color={group.url.length === 0 ? 'primary' : 'disabled'}
                    fontSize="small"
                  />
                </IconButton>
              </Tooltip>

              {group.url.length && (updating || updatingAll) ? (
                <CircularProgress color="secondary" />
              ) : (
                <Tooltip title="刷新">
                  <IconButton
                    onClick={this.handleGroupUpdate}
                    onFocus={event => {
                      event.stopPropagation();
                    }}
                    disabled={!group.url.length || updating || updatingAll}
                  >
                    <RefreshIcon
                      color={group.url.length === 0 ? 'disabled' : 'primary'}
                      fontSize="small"
                    />
                  </IconButton>
                </Tooltip>
              )}

              <Tooltip title="编辑">
                <IconButton
                  onClick={this.groupEditDialogOpen}
                  onFocus={event => {
                    event.stopPropagation();
                  }}
                >
                  <EditIcon color="primary" fontSize="small" />
                </IconButton>
              </Tooltip>
              <Tooltip title="删除">
                <IconButton
                  onClick={e => {
                    e.stopPropagation();
                    this.setState({ groupDeleteDialogOpen: true });
                  }}
                  onFocus={event => {
                    event.stopPropagation();
                  }}
                >
                  <DeleteIcon color="error" fontSize="small" />
                </IconButton>
              </Tooltip>
            </div>
            Last Update: {group.lastUpdate}
          </Typography>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails>
          <List className={styles.nodes}>
            {Array.isArray(group.nodes) &&
              group.nodes.map(node => (
                <>
                  <ListItem
                    button={true}
                    onClick={() => {
                      handleNodeClick(node);
                    }}
                  >
                    <ListItemText>
                      <Chip label={node.nodeType} color="secondary" clickable={true} />
                      <Chip label={node.remarks || node.ps || '未命名'} color="primary" clickable={true} />
                    </ListItemText>
                    <ListItemSecondaryAction>
                      <Button
                        disabled={group.url.length > 0}
                        onClick={event => {
                          this.nodeEditDialogOpen(event, node);
                        }}
                      >
                        编辑
                      </Button>
                      <Button
                        color="secondary"
                        onClick={e => {
                          e.stopPropagation();
                          this.setState({
                            nodeDeleteDialogOpen: true,
                            deletedNode: node,
                          });
                        }}
                      >
                        删除
                      </Button>
                    </ListItemSecondaryAction>
                  </ListItem>
                </>
              ))}
          </List>
        </ExpansionPanelDetails>
        <NodeEditDialog
          open={nodeEditDialogOpen}
          dialogClose={this.nodeEditDialogClose}
          node={editedNode}
          handleNodeAdd={this.handleNodeAdd}
          handleNodeEdit={this.handleNodeEdit}
        />
        <Dialog
          open={nodeDeleteDialogOpen}
          onClose={() => {
            this.setState({ nodeDeleteDialogOpen: false });
          }}
        >
          <DialogTitle>确定删除?</DialogTitle>
          <DialogActions>
            <Button color="primary" autoFocus={true} onClick={this.nodeDeleteDialogClose}>
              取消
            </Button>
            <Button color="secondary" onClick={this.handleNodeDelete} disabled={nodeDeleting}>
              删除
            </Button>
          </DialogActions>
        </Dialog>
        <GroupEditDialog
          open={groupEditDialogOpen}
          group={group}
          dialogClose={this.groupEditDialogClose}
          handleGroupEdit={this.handleGroupEdit}
        />
        <Dialog open={groupDeleteDialogOpen} onClose={this.groupDeleteDialogClose}>
          <DialogTitle>确定删除?</DialogTitle>
          <DialogActions>
            <Button color="primary" autoFocus={true} onClick={this.groupDeleteDialogClose}>
              取消
            </Button>
            <Button color="secondary" disabled={groupDeleting} onClick={this.handleGroupDelete}>
              删除
            </Button>
          </DialogActions>
        </Dialog>
      </ExpansionPanel>
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
      nodeDeleting: loading.effects['groups/deleteNode'],
      groupDeleting: loading.effects['groups/deleteGroup'],
    }),
  )(NodeGroup),
);
