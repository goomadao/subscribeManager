import React from 'react';
import { Dispatch } from 'redux';
import styles from './ProxyGroup.css';
import { makeStyles, withStyles } from '@material-ui/core/styles';
import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import MuiExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import Typography from '@material-ui/core/Typography';
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
  Dialog,
  DialogTitle,
  DialogActions,
  Chip,
  Tooltip,
  Divider,
} from '@material-ui/core';
import { Node, Group, ClashProxyGroupSelector } from '../typing';
import { connect } from 'dva';
import { withSnackbar, WithSnackbarProps } from 'notistack';
import SelectorEditDialog from './SelectorEditDialog';

const ExpansionPanelSummary = withStyles({
  content: {
    width: '100%',
    justifyContent: 'space-between',
  },
})(MuiExpansionPanelSummary);

export interface GroupProps extends WithSnackbarProps {
  selector: ClashProxyGroupSelector;
  selectors: ClashProxyGroupSelector[];
  groups: Group[];
  updatingAll: boolean;
  dispatch: Dispatch;
  selectorDeleting: boolean;
  handleNodeClick: (node: Node) => void;
}

export interface GroupState {
  updating: boolean;
  groupAddDialogOpen: boolean;
  selectorEditDialogOpen: boolean;
  selectorDeleteDialogOpen: boolean;
  deletedNode?: Node;
}

class NodeGroup extends React.Component<GroupProps, GroupState> {
  constructor(props: GroupProps) {
    super(props);
    this.state = {
      updating: false,
      selectorDeleteDialogOpen: false,
      groupAddDialogOpen: false,
      selectorEditDialogOpen: false,
    };
  }

  handleSelectorUpdate = (event: any) => {
    const { selector, dispatch, enqueueSnackbar } = this.props;
    event.stopPropagation();
    this.setState({
      updating: true,
    });
    dispatch({
      type: 'selectors/updateSelector',
      payload: {
        name: selector.name,
        type: selector.type,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(selector.name + '更新成功', {
            variant: 'success',
          });
        } else {
          enqueueSnackbar(msg || selector.name + '更新失败', {
            variant: 'error',
          });
        }
        this.setState({
          updating: false,
        });
      },
    });
  };

  handleSelectorEdit = (newSelector: ClashProxyGroupSelector) => {
    const { selector, dispatch, enqueueSnackbar } = this.props;
    return dispatch({
      type: 'selectors/editSelector',
      payload: {
        name: selector.name,
        selector: newSelector,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(selector.name + '修改成功', { variant: 'success' });
          this.groupEditDialogClose();
        } else {
          enqueueSnackbar(msg || selector.name + '修改失败', { variant: 'error' });
        }
      },
    });
  };

  handleSelectorDelete = (e: any) => {
    e.stopPropagation();
    const { selector, dispatch, enqueueSnackbar } = this.props;
    dispatch({
      type: 'selectors/deleteSelector',
      payload: {
        name: selector.name,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(selector.name + '删除成功', {
            variant: 'success',
          });
          this.setState({
            selectorDeleteDialogOpen: false,
          });
        } else {
          enqueueSnackbar(msg || selector.name + '删除失败', {
            variant: 'error',
          });
        }
      },
    });
  };

  selectorEditDialogOpen = (e: any) => {
    e.stopPropagation();
    this.setState({ selectorEditDialogOpen: true });
  };

  groupDeleteDialogClose = () => {
    this.setState({ selectorDeleteDialogOpen: false });
  };

  groupEditDialogClose = () => {
    this.setState({ selectorEditDialogOpen: false });
  };

  render() {
    const {
      selector,
      selectors,
      groups,
      selectorDeleting,
      updatingAll,
      handleNodeClick,
    } = this.props;
    const { updating, selectorDeleteDialogOpen, selectorEditDialogOpen } = this.state;
    return (
      <ExpansionPanel>
        <ExpansionPanelSummary
          aria-label="Expand"
          aria-controls="additional-actions1-content"
          id="additional-actions1-header"
        >
          <Typography color="textPrimary">
            {selector.name}
            <Chip
              label={selector.type}
              color="secondary"
              variant="outlined"
              style={{ marginLeft: '10px' }}
            />
          </Typography>
          <Typography color="textSecondary">
            <div className={styles.groupButtons}>
              {updating || updatingAll ? (
                <CircularProgress color="secondary" />
              ) : (
                <Tooltip title="刷新">
                  <IconButton
                    onClick={this.handleSelectorUpdate}
                    onFocus={event => {
                      event.stopPropagation();
                    }}
                    // disabled={!selector.url.length || updating || updatingAll}
                  >
                    <RefreshIcon
                      // color={selector.url.length === 0 ? 'disabled' : 'primary'}
                      color="primary"
                      fontSize="small"
                    />
                  </IconButton>
                </Tooltip>
              )}

              <Tooltip title="编辑">
                <IconButton
                  onClick={this.selectorEditDialogOpen}
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
                    this.setState({ selectorDeleteDialogOpen: true });
                  }}
                  onFocus={event => {
                    event.stopPropagation();
                  }}
                >
                  <DeleteIcon color="error" fontSize="small" />
                </IconButton>
              </Tooltip>
            </div>
          </Typography>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails>
          <div className={styles.selector}>
            <ExpansionPanel>
              <ExpansionPanelSummary>
                <Typography>代理组</Typography>
              </ExpansionPanelSummary>
              <ExpansionPanelDetails>
                <List className={styles.nodes}>
                  {selector.proxyGroups?.map((proxyGroup, i) => (
                    <>
                      <ListItem>
                        <ListItemText>{proxyGroup}</ListItemText>
                      </ListItem>
                      {i !== selector.proxyGroups.length - 1 && (
                        <Divider variant="middle" component="li" />
                      )}
                    </>
                  ))}
                </List>
              </ExpansionPanelDetails>
            </ExpansionPanel>
            <ExpansionPanel>
              <ExpansionPanelSummary>
                <Typography>节点</Typography>
              </ExpansionPanelSummary>
              <ExpansionPanelDetails>
                <List className={styles.nodes}>
                  {selector.proxies?.map(proxy => (
                    <>
                      <ListItem
                        button={true}
                        onClick={() => {
                          handleNodeClick(proxy);
                        }}
                      >
                        <ListItemText>
                          <Chip label={proxy.nodeType} color="secondary" clickable={true} />
                          <Chip
                            label={proxy.remarks || proxy.ps || '未命名'}
                            color="primary"
                            clickable={true}
                          />
                        </ListItemText>
                      </ListItem>
                    </>
                  ))}
                </List>
              </ExpansionPanelDetails>
            </ExpansionPanel>
          </div>
        </ExpansionPanelDetails>
        <SelectorEditDialog
          open={selectorEditDialogOpen}
          selector={selector}
          selectors={selectors}
          groups={groups}
          dialogClose={this.groupEditDialogClose}
          handleSelectorEdit={this.handleSelectorEdit}
        />
        <Dialog open={selectorDeleteDialogOpen} onClose={this.groupDeleteDialogClose}>
          <DialogTitle>确定删除?</DialogTitle>
          <DialogActions>
            <Button color="primary" autoFocus={true} onClick={this.groupDeleteDialogClose}>
              取消
            </Button>
            <Button
              color="secondary"
              disabled={selectorDeleting}
              onClick={this.handleSelectorDelete}
            >
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
      selectors,
      loading,
    }: {
      selectors: { selectors: ClashProxyGroupSelector[]; groups: Group[] };
      loading: {
        effects: { [key: string]: boolean };
      };
    }) => ({
      selectors: selectors.selectors,
      groups: selectors.groups,
      selectorDeleting: loading.effects['selectors/deleteSelector'],
    }),
  )(NodeGroup),
);
