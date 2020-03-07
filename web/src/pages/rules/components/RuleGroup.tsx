import React from 'react';
import styles from './RuleGroup.css';
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
} from '@material-ui/core';
import { RuleGroup } from '../typing';
import { connect, Dispatch } from 'dva';
import { withSnackbar, WithSnackbarProps } from 'notistack';
import RuleEditDialog from './RuleEditDialog';
import DisplayObject from '@/utils/DisplayObject';

const ExpansionPanelSummary = withStyles({
  content: {
    justifyContent: 'space-between',
  },
})(MuiExpansionPanelSummary);

export interface GroupProps extends WithSnackbarProps {
  group: RuleGroup;
  updatingAll: boolean;
  dispatch: Dispatch;
  ruleDeleting: boolean;
}

export interface GroupState {
  updating: boolean;
  groupAddDialogOpen: boolean;
  ruleEditDialogOpen: boolean;
  ruleDeleteDialogOpen: boolean;
}

class NodeGroup extends React.Component<GroupProps, GroupState> {
  constructor(props: GroupProps) {
    super(props);
    this.state = {
      updating: false,
      ruleDeleteDialogOpen: false,
      groupAddDialogOpen: false,
      ruleEditDialogOpen: false,
    };
  }

  handleRuleUpdate = (event: any) => {
    const { group, dispatch, enqueueSnackbar } = this.props;
    event.stopPropagation();
    this.setState({
      updating: true,
    });
    dispatch({
      type: 'rules/updateRule',
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

  handleruleEdit = (newRule: RuleGroup) => {
    const { group, dispatch, enqueueSnackbar } = this.props;
    return dispatch({
      type: 'rules/editRule',
      payload: {
        name: group.name,
        rule: newRule,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(group.name + '修改成功', { variant: 'success' });
          this.ruleEditDialogClose();
        } else {
          enqueueSnackbar(msg || group.name + '修改失败', { variant: 'error' });
        }
      },
    });
  };

  handleRuleDelete = (e: any) => {
    e.stopPropagation();
    const { group, dispatch, enqueueSnackbar } = this.props;
    dispatch({
      type: 'rules/deleteRule',
      payload: {
        name: group.name,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(group.name + '删除成功', {
            variant: 'success',
          });
          this.setState({
            ruleDeleteDialogOpen: false,
          });
        } else {
          enqueueSnackbar(msg || group.name + '删除失败', {
            variant: 'error',
          });
        }
      },
    });
  };

  ruleEditDialogOpen = (e: any) => {
    e.stopPropagation();
    this.setState({ ruleEditDialogOpen: true });
  };

  ruleDeleteDialogClose = () => {
    this.setState({ ruleDeleteDialogOpen: false });
  };

  ruleEditDialogClose = () => {
    this.setState({ ruleEditDialogOpen: false });
  };

  render() {
    const { group, ruleDeleting, updatingAll } = this.props;
    const { updating, ruleDeleteDialogOpen, ruleEditDialogOpen } = this.state;
    return (
      <ExpansionPanel>
        <ExpansionPanelSummary
          aria-label="Expand"
          aria-controls="additional-actions1-content"
          id="additional-actions1-header"
        >
          <Typography color="textPrimary">
            <div style={{ marginBottom: '10px' }}>
              {group.name}
              {group.url && (
                <Chip
                  label={group.url}
                  color="secondary"
                  variant="outlined"
                  style={{ marginLeft: '10px' }}
                />
              )}
            </div>
            <div>
              <Chip label={group.proxyGroup} color="primary" variant="outlined" />
            </div>
          </Typography>
          <Typography color="textSecondary">
            <div className={styles.groupButtons}>
              {group.url && group.url.length && (updating || updatingAll) ? (
                <CircularProgress color="secondary" />
              ) : (
                <Tooltip title="刷新">
                  <IconButton
                    onClick={this.handleRuleUpdate}
                    onFocus={event => {
                      event.stopPropagation();
                    }}
                    disabled={!group.url?.length || updating || updatingAll}
                  >
                    <RefreshIcon
                      color={!group.url?.length ? 'disabled' : 'primary'}
                      fontSize="small"
                    />
                  </IconButton>
                </Tooltip>
              )}

              <Tooltip title="编辑">
                <IconButton
                  onClick={this.ruleEditDialogOpen}
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
                    this.setState({ ruleDeleteDialogOpen: true });
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
            {group.customRules?.length > 0 &&
              group.customRules.map(rule => {
                let str = rule.split(',');
                return (
                  <>
                    <ListItem>
                      <ListItemText>
                        <Chip label={str[0]} color="secondary" />
                        {str.length > 1 && <Chip label={str[1]} color="primary" />}
                        <Chip label={group.proxyGroup} />
                      </ListItemText>
                    </ListItem>
                  </>
                );
              })}
            {group.rules?.length > 0 &&
              group.rules.map(rule => {
                let str = rule.split(',');
                return (
                  <ListItem>
                    <ListItemText>
                      <Chip label={str[0]} color="secondary" />
                      {str.length > 1 && <Chip label={str[1]} color="primary" />}
                      <Chip label={group.proxyGroup} />
                    </ListItemText>
                  </ListItem>
                );
              })}
          </List>
        </ExpansionPanelDetails>
        <RuleEditDialog
          open={ruleEditDialogOpen}
          group={group}
          dialogClose={this.ruleEditDialogClose}
          handleRuleEdit={this.handleruleEdit}
        />
        <Dialog open={ruleDeleteDialogOpen} onClose={this.ruleDeleteDialogClose}>
          <DialogTitle>确定删除?</DialogTitle>
          <DialogActions>
            <Button color="primary" autoFocus={true} onClick={this.ruleDeleteDialogClose}>
              取消
            </Button>
            <Button color="secondary" disabled={ruleDeleting} onClick={this.handleRuleDelete}>
              删除
            </Button>
          </DialogActions>
          {/* <DisplayObject {...this.props.loading} /> */}
        </Dialog>
      </ExpansionPanel>
    );
  }
}

export default withSnackbar(
  connect(
    ({
      rules,
      loading,
    }: {
      rules: RuleGroup[];
      loading: {
        effects: { [key: string]: boolean };
      };
    }) => ({
      groups: rules,
      ruleDeleting: loading.effects['rules/deleteRule'],
    }),
  )(NodeGroup),
);
